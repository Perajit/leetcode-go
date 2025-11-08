package generateparentheses

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}

	return generate(n, n, []string{""})
}

func generate(remainingOpen int, remainingClose int, initial []string) []string {
	if remainingOpen == 0 && remainingClose == 0 {
		return initial
	}

	N := len(initial)

	out := []string{}

	// add open bracket
	if remainingOpen > 0 {
		openApplied := make([]string, N)
		for i := 0; i < N; i++ {
			openApplied[i] = initial[i] + "("
		}

		openApplied = generate(remainingOpen-1, remainingClose, openApplied)
		out = append(out, openApplied...)
	}

	// add close bracket
	if remainingClose > 0 && remainingOpen < remainingClose {
		closeApplied := make([]string, N)

		for i := 0; i < N; i++ {
			closeApplied[i] = initial[i] + ")"
		}

		closeApplied = generate(remainingOpen, remainingClose-1, closeApplied)
		out = append(out, closeApplied...)
	}

	return out
}

/*
n=2
generate(2,2,[""])
->  generate(1,2,["("])														=>	["(())","()()"]
		->	generate(0,2,["(("])												=>	["(())"]
				->	generate(0,1,["(()"])
						->	generate(0,0,["(())"])
		->	generate(1,1,["()"])												=>	["()()"]
				->	generate(0,1,["()("])
						->	generate(0,0,["()()"])

n=3
generate(3,3,[""])
	->	generate(2,3,["("])													=>	["((()))","(()())","(())()","()(())","()()()"]
			->	generate(1,3,["(("])											=>	["((()))","(()())","(())()"]
					->	generate(0,3,["((("])										=>	["((()))"]
							->	generate(0,2,["((()"])
									->	generate(0,1,["((())"])
											-> generate(0,0,["((()))"])
					->	generate(1,2,["(()"])										=>	["(()())","(())()"]
							->	generate(0,2,["(()("])									=>	["(()())"]
									->	generate(0,1,["(()()"])
											->	generate(0,0,["(()())"])
							->	generate(1,1,["(())"])									=>	["(())()"]
									->	generate(0,1,["(())("])
											->	generate(0,0,["(())()"])
			->	generate(2,2,["()"])											=>	["()(())","()()()"]
					->	generate(1,2,["()("])										=>	["()(())","()()()"]
							->	generate(0,2,["()(("])								=>	["()(())"]
									->	generate(0,1,["()(()"])
											->	generate(0,0,["()(())"])
							->	generate(1,1,["()()"])								=>	["()()()"]
									->	generate(0,1,["()()("])
											->	generate(0,0,["()()()"])
*/
