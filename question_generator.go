package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io"
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
	desc = regexp.MustCompile("</?[a-zA-Z]+(?: [a-zA-Z_]+=\"[^\"]+\")*>").ReplaceAllString(desc, "")
	desc = html.UnescapeString(desc)
	desc = strings.TrimFunc(desc, func(r rune) bool {
		return r == ' ' || r == '\r' || r == '\n'
	})
	println(desc)
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
	var methodName, typeName, testMethodName string
	isClass := strings.Count(q.Code, "type ") > 0 && strings.Count(q.Code, "func ") > 1
	if isClass {
		fmt.Printf("%s is a class question.\n", q.Title)

		typeNameMatch := regexp.MustCompile("type (.+) struct {").FindAllStringSubmatch(q.Code, -1)
		typeName = typeNameMatch[0][1]
		testMethodName = "Test" + typeName
	} else {
		methodNameMatch := regexp.MustCompile("func (.+)\\(").FindAllStringSubmatch(q.Code, -1)
		methodName = methodNameMatch[0][1]
		testMethodName = "Test" + strings.ToUpper(methodName[0:1]) + methodName[1:]
	}

	desc := "// " + strings.ReplaceAll(q.Text, "\n", "\n// ")

	dirName := "./src/" + q.PackageName
	os.MkdirAll(dirName, os.ModePerm)

	questionFile := dirName + "/question.go"
	if info, _ := os.Stat(questionFile); info != nil {
		println("question.go file exist, skip.")
	} else if file, err := os.Create(questionFile); err != nil {
		panic(fmt.Sprint("Error while create question.go: ", err))
	} else {
		defer file.Close()
		if isClass {
			convertTemplate(file, "class_question_template", map[string]interface{}{
				"q":        q,
				"desc":     desc,
				"TypeName": typeName,
			})
		} else {
			convertTemplate(file, "method_question_template", map[string]interface{}{
				"q":          q,
				"desc":       desc,
				"MethodName": methodName,
			})
		}
	}

	answerFile := dirName + "/answer_test.go"
	if info, _ := os.Stat(answerFile); info != nil {
		println("answer_test.go file exist, skip.")
	} else if file, err := os.Create(answerFile); err != nil {
		panic(fmt.Sprint("Error while create answer_test.go: ", err))
	} else {
		defer file.Close()
		if isClass {
			convertTemplate(file, "class_answer_template", map[string]interface{}{
				"q":              q,
				"TestMethodName": testMethodName,
			})
		} else {
			convertTemplate(file, "method_answer_template", map[string]interface{}{
				"q":              q,
				"TestMethodName": testMethodName,
			})
		}
	}
}

func convertTemplate(writer io.Writer, templateName string, variables map[string]interface{}) {
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
	tmpl.Execute(writer, variables)
}
