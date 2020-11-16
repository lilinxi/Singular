package Singular

func GaussElimination(A, b Matrix) (x Matrix) {
	// step1 消元
	for i := 0; i < A.Rows(); i++ { // 选择消去的主元
		if A.Get(i, i) == 0 {
			panic("")
		}
		for j := i + 1; j < A.Rows(); j++ { // 当前消去需要更新的行
			m := A.Get(j, i) / A.Get(i, i)
			for k := i; k < A.Cols(); k++ { // 当前消去需要更新的列
				A.SetAdd(j, k, -m*A.Get(i, k))
			}
			b.SetAdd(j, 0, -m*b.Get(i, 0))
		}
	}

	// step2 回代
	x = b.Copy()
	for i := A.Rows() - 1; i >= 0; i-- {
		for j := i + 1; j < A.Rows(); j++ {
			x.SetAdd(i, 0, -A.Get(i, j)*x.Get(j, 0))
		}
		x.SetScale(i, 0, 1/A.Get(i, i))
	}

	//fmt.Println(A)
	//fmt.Println(b)
	//fmt.Println(x)

	return x
}
