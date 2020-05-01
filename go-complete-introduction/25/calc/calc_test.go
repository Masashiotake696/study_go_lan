package calc

// 通常のテスト
// func TestCompute(t *testing.T) {
// 	s, err := Compute("1+1")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	if s != "2" {
// 		t.Errorf("Compute(1+1) = %s, want 2", s)
// 	}
// }

// テーブル駆動テスト
// func TestCompute(t *testing.T) {
// 	computeTests := []struct {
// 		in  string
// 		out string
// 	}{
// 		{"1+1", "2"},
// 		{"1.0/2.0", "0.5"},
// 	}

// 	for _, test := range computeTests {
// 		s, err := Compute(test.in)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		if s != test.out {
// 			t.Errorf("Compute(%s) = %s, want %s", test.in, s, test.out)
// 		}
// 	}
// }

// ブラックボックステスト
// func TestCompute(t *testing.T) {
// 	add := func(a, b int16) bool {
// 		s, err := Compute(fmt.Sprintf("%d+%d", a, b))
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		expected := strconv.Itoa(int(a) + int(b))

// 		if s != expected {
// 			t.Logf("Compute(%d+%d) = %s, want %s", a, b, s, expected)
// 			return false
// 		}
// 		return true
// 	}

// 	if err := quick.Check(add, nil); err != nil {
// 		t.Fatal(err)
// 	}
// }

// サブテスト
// func TestCompute(t *testing.T) {
// 	t.Run("add sub", func(t *testing.T) {
// 		testCompute(t, "1+1", "2")
// 		testCompute(t, "-2+1", "-1")
// 	})
// 	t.Run("div", func(t *testing.T) {
// 		testCompute(t, "1.0/2.0", "0.5")
// 		testCompute(t, "2.0/1.0", "2")
// 	})
// }

// サブテストとテーブル駆動テスト
// func TestCompute(t *testing.T) {
// 	cases := []struct {
// 		name string
// 		in   string
// 		out  string
// 	}{
// 		{"add", "1+1", "2"},
// 		{"sub", "-2+1", "-1"},
// 		{"div", "1.0/2.0", "0.5"},
// 	}

// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			testCompute(t, c.in, c.out)
// 		})
// 	}
// }

// サブテスト用関数
// func testCompute(t *testing.T, in, expected string) {
// 	s, err := Compute(in)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if s != expected {
// 		t.Errorf("Compute(%s) = %s, want %s", in, s, expected)
// 	}
// }

// ヘルパーを使ったテスト
// func TestCompute(t *testing.T) {
// 	t.Run("add sub", func(t *testing.T) {
// 		testCompute(t, "1+1", "2")
// 		testCompute(t, "-2+1", "-1")
// 		testCompute(t, "1+1", "二")
// 	})
// }

// ヘルパーを使ったテスト用関数
// func testCompute(t *testing.T, in, expected string) {
// 	t.Helper()
// 	s, err := Compute(in)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if s != expected {
// 		t.Errorf("Compute(%s) = %s, want %s", in, s, expected)
// 	}
// }
