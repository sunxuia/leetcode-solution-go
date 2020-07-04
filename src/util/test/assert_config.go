package test

// assert 的配置
type assertConfig struct {
	shouldChecks         *matchTree
	hasOrders            *matchTree
	validators           *matchTree
	expectTypeValidators map[string]func(*Handler)
	actualTypeValidators map[string]func(*Handler)
	numberTolerance      float64
}

func NewAssertConfig() *assertConfig {
	res := &assertConfig{
		shouldChecks:         newMatchTree(),
		hasOrders:            newMatchTree(),
		validators:           newMatchTree(),
		expectTypeValidators: map[string]func(*Handler){},
		actualTypeValidators: map[string]func(*Handler){},
	}
	res.shouldChecks.setNodeValue("**", true)
	res.hasOrders.setNodeValue("**", true)
	return res
}

func (cf *assertConfig) Check(fieldName string) *assertConfig {
	cf.shouldChecks.setNodeValue(fieldName, true)
	return cf
}

func (cf *assertConfig) UnCheck(fieldName string) *assertConfig {
	cf.shouldChecks.setNodeValue(fieldName, false)
	return cf
}

func (cf *assertConfig) Order(fieldName string) *assertConfig {
	cf.hasOrders.setNodeValue(fieldName, true)
	return cf
}

func (cf *assertConfig) UnOrder(fieldName string) *assertConfig {
	cf.hasOrders.setNodeValue(fieldName, false)
	return cf
}

func (cf *assertConfig) AddValidator(fieldName string, validateMethod func(*Handler)) *assertConfig {
	cf.validators.setNodeValue(fieldName, validateMethod)
	return cf
}

func (cf *assertConfig) AddExpectTypeValidator(typeName string, validateMethod func(*Handler)) {
	cf.expectTypeValidators[typeName] = validateMethod
}

func (cf *assertConfig) AddActualTypeValidator(typeName string, validateMethod func(*Handler)) {
	cf.actualTypeValidators[typeName] = validateMethod
}

func (cf *assertConfig) FloatPrecision(tolerance float64) *assertConfig {
	if tolerance < 0 {
		panic("Tolerance must >= 0.")
	}
	cf.numberTolerance = tolerance
	return cf
}
