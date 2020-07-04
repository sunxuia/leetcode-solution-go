package test

import (
	"fmt"
	"regexp"
	"runtime/debug"
	"strings"
	"testing"
)

// testHelper 断言工具
type testHelper struct {
	T            *testing.T
	caseNames    []string
	assertConfig *assertConfig
}

var answerFilePattern = regexp.MustCompile(`/src/q\d\d\d_[^/]+/q\d\d\d_answer_.+\.go:\d+ \+0x[0-9a-fA-F]+`)

var questionFilePattern = regexp.MustCompile(`/src/q\d\d\d_[^/]+/q\d\d\d_question\.go:\d+ \+0x[0-9a-fA-F]+`)

var testFilePattern = regexp.MustCompile(`/src/util/test/.+\.go:\d+ \+0x[0-9a-fA-F]+$`)

// NewTestHelper 断言工具构造函数
func NewTestHelper(t *testing.T) *testHelper {
	res := &testHelper{
		T:            t,
		caseNames:    make([]string, 0),
		assertConfig: defaultAssertConfig,
	}
	return res
}

var defaultAssertConfig = NewAssertConfig()

func (th *testHelper) AssertConfig(config *assertConfig) *testHelper {
	th.assertConfig = config
	return th
}

func (th *testHelper) TestCase(name string) func() {
	stack := string(debug.Stack())
	lines := strings.Split(stack, "\n")
	var line string
	for i := 5; i < len(lines); i++ {
		if questionFilePattern.MatchString(lines[i]) {
			line = lines[i]
			break
		}
	}
	if len(line) > 0 {
		name += "\n(line: " + strings.Trim(line, "\t ") + ")"
	}

	th.caseNames = append(th.caseNames, name)

	return func() {
		if r := recover(); r != nil {
			th.fail(r)
		}
		th.caseNames = th.caseNames[:len(th.caseNames)-1]
	}
}

func (th *testHelper) fail(r interface{}) {
	sb := strings.Builder{}
	if e, ok := r.(*AssertError); ok {
		sb.WriteString("AssertError!\n")
		sb.WriteString(e.String())
	} else {
		var caseName string
		if len(th.caseNames) > 0 {
			caseName = th.caseNames[len(th.caseNames)-1]
		}
		stacks := string(debug.Stack())
		lines := strings.Split(stacks, "\n")[3:]
		if idxes := answerFilePattern.FindAllStringIndex(stacks, -1); len(idxes) > 1 {
			sb.WriteString("Answer Code Error!")
			if len(caseName) > 0 {
				sb.WriteString(" With test case " + caseName)
			}
			sb.WriteByte('\n')
			sb.WriteString(fmt.Sprint(r, "\n"))
			i := 1
			for ; i < len(lines); i += 2 {
				if !testFilePattern.MatchString(lines[i]) {
					break
				}
			}
			if i < len(lines) && strings.Contains(lines[i], "/src/runtime/panic.go:") {
				i += 2
			}
			for ; i < len(lines); i++ {
				sb.WriteString(strings.Trim(lines[i], "\t "))
				sb.WriteByte('\n')
			}
		} else {
			sb.WriteString("Unexpected Error!")
			if len(caseName) > 0 {
				sb.WriteString(" With test case " + caseName)
			}
			sb.WriteByte('\n')
			sb.WriteString(fmt.Sprint(r, "\n"))
			for i := 0; i < len(lines); i++ {
				sb.WriteString(strings.Trim(lines[i], "\t "))
				sb.WriteByte('\n')
			}
		}
	}

	th.T.Error(sb.String())
	th.T.FailNow()
}

func (th *testHelper) NewTestCase() func() {
	var name string
	switch len(th.caseNames) {
	case 0:
		name = "1st test case"
	case 1:
		name = "2nd test case"
	default:
		name = fmt.Sprint(len(th.caseNames)+1, "th test case")
	}

	return th.TestCase(name)
}

func (th *testHelper) Expect(val interface{}) *dataExpect {
	return &dataExpect{
		th:      th,
		expects: []interface{}{val},
		config:  th.assertConfig,
	}
}

func (th *testHelper) AssertEqual(expect, actual interface{}) {
	th.Expect(expect).Assert(actual)
}
