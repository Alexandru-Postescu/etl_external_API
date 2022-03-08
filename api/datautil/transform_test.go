package datautil

import (
	"testing"

	"example.com/go-api/model"
	"github.com/google/go-cmp/cmp"
)

// TODO: failing test, run it with -count=1 to clear cache
func TestTransform(t *testing.T) {
	dataPreparator := NewDataTransformer()
	type args struct {
		input []model.Person
	}
	name1 := "test1"
	name2 := "test2"

	person1 := model.Person{
		"alex",
		"poste",
		"idk",
		"idk",
		"idk",
		"idk",
	}
	personIdentic := model.Person{
		"alex",
		"poste",
		"idk",
		"idk",
		"idk",
		"idk",
	}
	person2 := model.Person{
		"aalex",
		"poste",
		"idk",
		"idk",
		"idk",
		"idk",
	}

	person3 := model.Person{
		"bblex",
		"poste",
		"idk",
		"idk",
		"idk",
		"idk",
	}
	person4 := model.Person{
		"blex",
		"poste",
		"idk",
		"idk",
		"idk",
		"idk",
	}
	input1 := []model.Person{person1, person2, personIdentic}
	input2 := []model.Person{person3, person4}
	want1 := map[string][]model.Person{
		"a": []model.Person{person1, person2},
	}
	want2 := map[string][]model.Person{
		"b": input2,
	}

	tests := []struct {
		name    string
		args    args
		want    map[string][]model.Person
		wantErr bool
	}{
		{name1, args{
			input: input1,
		}, want1, false},
		{name2, args{
			input: input2,
		}, want2, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dataPreparator.Transform(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transform() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Transform() = %v, want %v", got, tt.want)
			}
		})
	}
}
