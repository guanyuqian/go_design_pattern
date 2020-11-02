package _8_Filter_Pattern

//步骤 1
//创建一个类，在该类上应用标准。
type Person struct {
	name, gender, maritalStatus string
}

//步骤 2
//为标准（Criteria）创建一个接口。
type Criteria interface {
	meetCriteria([]Person) []Person
}

//步骤 3
//创建实现了 Criteria 接口的实体类。
type CriteriaMale struct{}

func (receiver CriteriaMale) meetCriteria(persons []Person) (res []Person) {
	for _, person := range persons {
		if person.gender == "Male" {
			res = append(res, person)
		}
	}
	return
}

type CriteriaFemale struct{}

func (receiver CriteriaFemale) meetCriteria(persons []Person) (res []Person) {
	for _, person := range persons {
		if person.gender == "Female" {
			res = append(res, person)
		}
	}
	return
}

type CriteriaSingle struct{}

func (receiver CriteriaSingle) meetCriteria(persons []Person) (res []Person) {
	for _, person := range persons {
		if person.maritalStatus == "Single" {
			res = append(res, person)
		}
	}
	return
}

type AndCriteria struct {
	criteria, otherCriteria Criteria
}

func (receiver AndCriteria) meetCriteria(persons []Person) []Person {
	firstCriteriaItems := receiver.criteria.meetCriteria(persons)
	return receiver.otherCriteria.meetCriteria(firstCriteriaItems)
}

type OrCriteria struct {
	criteria, otherCriteria Criteria
}

func (receiver OrCriteria) meetCriteria(persons []Person) []Person {
	firstCriteriaItems := receiver.criteria.meetCriteria(persons)
	otherCriteriaItems := receiver.otherCriteria.meetCriteria(persons)
	for _, person := range otherCriteriaItems {
		exist := false
		for _, v := range firstCriteriaItems {
			if person == v {
				exist = true
				break
			}
		}
		if !exist {
			firstCriteriaItems = append(firstCriteriaItems, person)
		}
	}
	return firstCriteriaItems
}
