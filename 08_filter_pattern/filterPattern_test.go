package _8_filter_pattern

import (
	"reflect"
	"testing"
)

//步骤4
//使用不同的标准（Criteria）和它们的结合来过滤 Person 对象的列表。
var persons []Person

func init() {
	persons = []Person{
		{"Robert", "Male", "Single"},
		{"John", "Male", "Married"},
		{"Laura", "Female", "Married"},
		{"Diana", "Female", "Single"},
		{"Mike", "Male", "Single"},
		{"Bobby", "Male", "Single"},
	}
}

func TestAndCriteria_meetCriteria(t *testing.T) {
	type fields struct {
		criteria      Criteria
		otherCriteria Criteria
	}
	type args struct {
		persons []Person
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Person
	}{
		{"AndCriteria", fields{new(CriteriaMale), new(CriteriaSingle)}, args{persons},
			[]Person{
				{"Robert", "Male", "Single"},
				{"Mike", "Male", "Single"},
				{"Bobby", "Male", "Single"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := AndCriteria{
				criteria:      tt.fields.criteria,
				otherCriteria: tt.fields.otherCriteria,
			}
			got := receiver.meetCriteria(tt.args.persons)
			if len(got) != len(tt.want) {
				t.Errorf("meetCriteria() = %v, want %v", got, tt.want)
			}
			for _, v1 := range got {
				exist := false
				for _, v2 := range tt.want {
					if v1 == v2 {
						exist = true
						break
					}
				}
				if !exist {
					t.Errorf("meetCriteria() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestCriteriaFemale_meetCriteria(t *testing.T) {
	type args struct {
		persons []Person
	}
	tests := []struct {
		name    string
		args    args
		wantRes []Person
	}{
		{"CriteriaFemale", args{persons},
			[]Person{
				{"Laura", "Female", "Married"},
				{"Diana", "Female", "Single"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := CriteriaFemale{}
			if gotRes := receiver.meetCriteria(tt.args.persons); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("meetCriteria() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestCriteriaMale_meetCriteria(t *testing.T) {
	type args struct {
		persons []Person
	}
	tests := []struct {
		name    string
		args    args
		wantRes []Person
	}{
		{"CriteriaMale", args{persons},
			[]Person{
				{"Robert", "Male", "Single"},
				{"John", "Male", "Married"},
				{"Mike", "Male", "Single"},
				{"Bobby", "Male", "Single"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := CriteriaMale{}
			if gotRes := receiver.meetCriteria(tt.args.persons); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("meetCriteria() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestCriteriaSingle_meetCriteria(t *testing.T) {
	type args struct {
		persons []Person
	}
	tests := []struct {
		name    string
		args    args
		wantRes []Person
	}{
		// TODO: add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := CriteriaSingle{}
			if gotRes := receiver.meetCriteria(tt.args.persons); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("meetCriteria() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestOrCriteria_meetCriteria(t *testing.T) {
	type fields struct {
		criteria      Criteria
		otherCriteria Criteria
	}
	type args struct {
		persons []Person
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Person
	}{
		{"OrCriteria", fields{new(CriteriaMale), new(CriteriaFemale)}, args{persons},
			[]Person{{"Robert", "Male", "Single"},
				{"John", "Male", "Married"},
				{"Laura", "Female", "Married"},
				{"Diana", "Female", "Single"},
				{"Mike", "Male", "Single"},
				{"Bobby", "Male", "Single"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := OrCriteria{
				criteria:      tt.fields.criteria,
				otherCriteria: tt.fields.otherCriteria,
			}
			got := receiver.meetCriteria(tt.args.persons)
			if len(got) != len(tt.want) {
				t.Errorf("meetCriteria() = %v, want %v", got, tt.want)
			}
			for _, v1 := range got {
				exist := false
				for _, v2 := range tt.want {
					if v1 == v2 {
						exist = true
						break
					}
				}
				if !exist {
					t.Errorf("meetCriteria() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
