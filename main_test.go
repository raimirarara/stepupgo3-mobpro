package main

import "testing"

func TestGoVet(t *testing.T) {

	want := 
	`github.com/gostaticanalysis/skeleton/v2@v2.0.0
# github.com/gostaticanalysis/skeleton/v2/skeleton
{}
# github.com/gostaticanalysis/skeleton/v2
{}
# github.com/gostaticanalysis/skeleton/v2/skeleton_test
{}

github.com/gostaticanalysis/skeleton/v2@v2.0.1
# github.com/gostaticanalysis/skeleton/v2/skeleton
{}
# github.com/gostaticanalysis/skeleton/v2
{}
# github.com/gostaticanalysis/skeleton/v2/skeleton_test
{}

github.com/gostaticanalysis/skeleton/v2@v2.0.2
# github.com/gostaticanalysis/skeleton/v2/skeleton
{}
# github.com/gostaticanalysis/skeleton/v2
{}
# github.com/gostaticanalysis/skeleton/v2/skeleton_test
{}

github.com/gostaticanalysis/skeleton/v2@v2.0.3
# github.com/gostaticanalysis/skeleton/v2/skeleton
{}
# github.com/gostaticanalysis/skeleton/v2
{}
# github.com/gostaticanalysis/skeleton/v2/skeleton_test
{}

github.com/gostaticanalysis/skeleton/v2@v2.0.4
# github.com/gostaticanalysis/skeleton/v2/skeleton
{}
# github.com/gostaticanalysis/skeleton/v2
{}
# github.com/gostaticanalysis/skeleton/v2/skeleton_test
{}

`



	t.Run("test1" , func(t *testing.T) {
		got := goVet("github.com/gostaticanalysis/skeleton/v2")
		if got != want {
			t.Errorf("got: %s :", got)
		}
	})

	

}