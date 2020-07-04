package test

type dataExpect struct {
	th      *testHelper
	config  *assertConfig
	expects []interface{}
}

func (de *dataExpect) OrExpect(val interface{}) *dataExpect {
	de.expects = append(de.expects, val)
	return de
}

func (de *dataExpect) Config(config *assertConfig) *dataExpect {
	de.config = config
	return de
}

func (de *dataExpect) Assert(actual interface{}) {
	var err interface{}
	for _, expect := range de.expects {
		err = de.assertWithError(expect, actual)
		if err == nil {
			break
		}
	}

	if err != nil {
		inTestCase := len(de.th.caseNames) > 0
		if inTestCase {
			if len(de.expects) > 1 {
				handler := de.newHandler(de.expects, actual)
				handler.Error("None of expects matches actual.")
			} else {
				panic(err)
			}
		} else {
			de.th.fail(err)
		}
	}
}

func (de *dataExpect) assertWithError(expect, actual interface{}) (err interface{}) {
	defer func() {
		if r := recover(); r != nil {
			err = r
		}
	}()

	handler := de.newHandler(expect, actual)
	handler.Assert()
	return
}

func (de *dataExpect) newHandler(expect, actual interface{}) *Handler {
	res := &Handler{
		FieldName:  "",
		FullExpect: expect,
		FullActual: actual,
		config:     de.config,
	}
	if len(de.th.caseNames) > 0 {
		res.caseName = de.th.caseNames[len(de.th.caseNames)-1]
	}
	res.ChangeValue(expect, actual)
	return res
}
