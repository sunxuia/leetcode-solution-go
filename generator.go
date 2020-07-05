package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

func main() {
	questionID, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		panic(fmt.Sprint("Argument error: ", err))
	}

	question := getQuestion(questionID)
	generateFiles(question)
}

type question struct {
	ID          uint64
	Title       string
	TitleSlug   string
	PackageName string
	URL         string
	Difficulty  string
	PaidOnly    bool
	Text        string
	Code        string
}

type questionMethod struct {
	raw        string
	methodName string
	returnType string
	callerType string
	argType    []string
	argName    []string
}

func getQuestion(questionID uint64) *question {
	algorithms := requestAlgorithms()
	statStatusPairs := reflect.ValueOf(algorithms["stat_status_pairs"])
	var ques, stat map[string]interface{}
	for i := statStatusPairs.Len() - 1; i >= 0 && ques == nil; i-- {
		ques = statStatusPairs.Index(i).Interface().(map[string]interface{})
		stat = ques["stat"].(map[string]interface{})
		id := uint64(stat["frontend_question_id"].(float64))
		if questionID != id {
			ques = nil
		}
	}
	if ques == nil {
		panic(fmt.Sprint("Cannot find question ", questionID))
	}

	res := new(question)
	res.PaidOnly = ques["paid_only"].(bool)
	if res.PaidOnly {
		panic(fmt.Sprint("Question ", res.Title, " is paid only!"))
	}

	res.ID = questionID
	res.Title = stat["question__title"].(string)
	res.TitleSlug = stat["question__title_slug"].(string)
	res.PackageName = fmt.Sprintf("q%03d_%s", questionID, strings.ReplaceAll(res.TitleSlug, "-", "_"))
	res.URL = "https://leetcode.com/problems/" + res.TitleSlug + "/"
	difficulty := ques["difficulty"].(map[string]interface{})
	res.Difficulty = map[float64]string{
		1: "Easy",
		2: "Medium",
		3: "Hard",
	}[difficulty["level"].(float64)]

	quesionContent := requestQuestionContent(res.TitleSlug)
	question := quesionContent["question"].(map[string]interface{})

	desc := question["content"].(string)
	desc = html.UnescapeString(desc)
	desc = regexp.MustCompile("</?[a-zA-Z]+(?: [a-zA-Z_]+=\"[^\"]+\")*/?>").ReplaceAllString(desc, "")
	desc = html.UnescapeString(desc)
	desc = strings.TrimFunc(desc, func(r rune) bool {
		return r == ' ' || r == '\r' || r == '\n'
	})
	res.Text = desc

	codeSnippets := reflect.ValueOf(question["codeSnippets"])
	var codeStr string
	for i := codeSnippets.Len() - 1; i >= 0 && len(codeStr) == 0; i-- {
		code := codeSnippets.Index(i).Interface().(map[string]interface{})
		if "Go" == code["lang"].(string) {
			codeStr = code["code"].(string)
		}
	}
	if len(codeStr) == 0 {
		panic(fmt.Sprint("Question ", res.Title, " does not have go solution!"))
	}
	res.Code = codeStr

	return res
}

func requestAlgorithms() map[string]interface{} {
	resp, err := http.Get("https://leetcode.com/api/problems/algorithms/")
	if err != nil {
		panic(fmt.Sprint("Error while get question list ", err))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprint("Error while get question list ", err))
	}

	res := map[string]interface{}{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		panic(fmt.Sprint("Error while parse question list ", err))
	}
	return res
}

func requestQuestionContent(titleSlug string) map[string]interface{} {
	request := map[string]interface{}{
		"operationName": "questionData",
		"variables": map[string]interface{}{
			"titleSlug": titleSlug,
		},
		"query": `query questionData($titleSlug: String!) {
			question(titleSlug: $titleSlug) {
			  questionId
			  questionFrontendId
			  boundTopicId
			  title
			  titleSlug
			  content
			  translatedTitle
			  translatedContent
			  isPaidOnly
			  difficulty
			  likes
			  dislikes
			  isLiked
			  similarQuestions
			  contributors {
				username
				profileUrl
				avatarUrl
				__typename
			  }
			  langToValidPlayground
			  topicTags {
				name
				slug
				translatedName
				__typename
			  }
			  companyTagStats
			  codeSnippets {
				lang
				langSlug
				code
				__typename
			  }
			  stats
			  hints
			  solution {
				id
				canSeeDetail
				paidOnly
				__typename
			  }
			  status
			  sampleTestCase
			  metaData
			  judgerAvailable
			  judgeType
			  mysqlSchemas
			  enableRunCode
			  enableTestMode
			  enableDebugger
			  envInfo
			  libraryUrl
			  adminUrl
			  __typename
			}
		  }`,
	}
	requestBody, _ := json.Marshal(&request)
	resp, err := http.Post("https://leetcode.com/graphql", "application/json", bytes.NewReader(requestBody))
	if err != nil {
		panic(fmt.Sprint("Error while parse question ", titleSlug, " : ", err))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprint("Error while parse question ", titleSlug, " : ", err))
	}

	bodyMap := map[string]interface{}{}
	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		panic(fmt.Sprint("Error while parse question data ", err))
	}
	return bodyMap["data"].(map[string]interface{})
}

func generateFiles(q *question) {
	code, typeName, methods := handleCode(q.Code)
	if len(typeName) > 0 {
		fmt.Printf("%s is a class question.\n", q.Title)
		handleClassFile(code, q, typeName, methods)
	} else {
		handleMethodFile(code, q, methods[0])
	}
}

func handleCode(code string) (string, string, []*questionMethod) {
	code = regexp.MustCompile(`/\*[\w\W]*?\*/`).ReplaceAllString(code, "")
	code = regexp.MustCompile(`//.+\n`).ReplaceAllString(code, "")
	code = strings.Trim(code, "\n\t ")

	var typeName string
	structRegex := regexp.MustCompile(`type ([^ ]+) struct {`)
	if m := structRegex.FindAllStringSubmatch(code, -1); m != nil {
		typeName = m[0][1]
	}

	methodRegex := regexp.MustCompile(`func(?: \(([^)]+)\))? ([a-zA-Z0-9_]+)\s*\(([^)]*)\)(?: (.+))?\s*{`)
	argRegex := regexp.MustCompile("([^ ]+) ([^ ]+)")
	methods := make([]*questionMethod, 0)
	if m := methodRegex.FindAllStringSubmatchIndex(code, -1); m != nil {
		for i := 0; i < len(m); i++ {
			qm := &questionMethod{
				raw:        strings.Trim(code[m[i][0]:m[i][1]], " \t\n"),
				methodName: code[m[i][4]:m[i][5]],
				returnType: strings.Trim(code[m[i][8]:m[i][9]], " "),
				argName:    make([]string, 0),
				argType:    make([]string, 0),
			}
			if m[i][2] > -1 {
				qm.callerType = argRegex.FindAllStringSubmatch(code[m[i][2]:m[i][3]], -1)[0][2]
			}
			if m[i][6] > -1 && len(code[m[i][6]:m[i][7]]) > 0 {
				args := strings.Split(code[m[i][6]:m[i][7]], ",")
				for _, arg := range args {
					aMatch := argRegex.FindAllStringSubmatch(arg, -1)
					qm.argName = append(qm.argName, aMatch[0][1])
					qm.argType = append(qm.argType, aMatch[0][2])
				}
			}
			methods = append(methods, qm)
		}
	}
	return code, typeName, methods
}

func handleClassFile(code string, q *question, typeName string, methods []*questionMethod) {
	testMethodName := "Test" + typeName
	iMethods := make([]string, 0)
	for _, method := range methods {
		if len(method.callerType) > 0 {
			callerEnd := strings.Index(method.raw, ")") + 1
			iMethods = append(iMethods, method.raw[callerEnd:len(method.raw)-1])
		}
	}

	desc := "// " + strings.ReplaceAll(q.Text, "\n", "\n// ")

	dirName := "./src/" + q.PackageName
	_ = os.MkdirAll(dirName, os.ModePerm)

	questionFileName := fmt.Sprintf("%s/q%03d_%s", dirName, q.ID, "question.go")
	outputFile(questionFileName, "class_question_template", map[string]interface{}{
		"q":        q,
		"desc":     desc,
		"TypeName": typeName,
		"Methods":  iMethods,
	})

	answerFileName := fmt.Sprintf("%s/q%03d_%s", dirName, q.ID, "answer_test.go")
	outputFile(answerFileName, "class_answer_template", map[string]interface{}{
		"q":              q,
		"Code":           code,
		"TypeName":       typeName,
		"TestMethodName": testMethodName,
	})
}

func outputFile(fileName string, templateName string, variables map[string]interface{}) {
	if info, _ := os.Stat(fileName); info != nil {
		println(fileName + " file exist, skip.")
	} else if file, err := os.Create(fileName); err != nil {
		panic(fmt.Sprint("Error while create ", fileName, "\n", err))
	} else {
		defer func() {
			_ = file.Close()
		}()
		filePath := "./template/" + templateName + ".txt"
		bs, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(fmt.Sprint("Error while read tempalte ", templateName, " : ", err))
		}
		str := string(bs)

		tmpl, err := template.New("test").Parse(str)
		if err != nil {
			panic(err)
		}
		_ = tmpl.Execute(file, variables)
	}
}

func handleMethodFile(code string, q *question, method *questionMethod) {
	testMethodName := "Test" + strings.ToUpper(method.methodName[0:1]) + method.methodName[1:]

	funcSb := strings.Builder{}
	funcSb.WriteString("func(")
	for i, t := range method.argType {
		if i > 0 {
			funcSb.WriteString(", ")
		}
		funcSb.WriteString(t)
	}
	funcSb.WriteString(")")
	if len(method.returnType) > 0 {
		funcSb.WriteString(" ")
		funcSb.WriteString(method.returnType)
	}

	desc := "// " + strings.ReplaceAll(q.Text, "\n", "\n// ")

	dirName := "./src/" + q.PackageName
	_ = os.MkdirAll(dirName, os.ModePerm)

	questionFileName := fmt.Sprintf("%s/q%03d_%s", dirName, q.ID, "question.go")
	outputFile(questionFileName, "method_question_template", map[string]interface{}{
		"q":          q,
		"desc":       desc,
		"MethodName": method.methodName,
		"ReturnType": method.returnType,
		"Func":       funcSb.String(),
	})

	answerFileName := fmt.Sprintf("%s/q%03d_%s", dirName, q.ID, "answer_test.go")
	outputFile(answerFileName, "method_answer_template", map[string]interface{}{
		"q":              q,
		"Code":           code,
		"MethodName":     method.methodName,
		"TestMethodName": testMethodName,
	})
}
