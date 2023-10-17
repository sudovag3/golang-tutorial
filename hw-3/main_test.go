package main

import "testing"

func TestBestRepeats(t *testing.T) {
	// Создаем структуру с тестовыми данными
	tests := []struct {
		input    string
		expected []Word
	}{
		{
			input: "Ты очень красивый, я немного влюбилась. Как ты считаешь, я красивая? Если что, я свободна сегодня, можем встретиться, если ты не занят. Ты крутой",
			expected: []Word{
				{Name: "ты", Repeat: 4},
				{Name: "я", Repeat: 3},
			},
		},
		{
			input: "раз два два три три три четыре четыре четыре четыре",
			expected: []Word{
				{Name: "четыр", Repeat: 4},
				{Name: "три", Repeat: 3},
			},
		},
	}

	for _, test := range tests {
		result := BestRepeats(test.input)[:2]

		// Проверяем, что результат соответствует ожидаемому
		if len(result) != len(test.expected) {
			t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
			continue
		}

		for i := range result {
			if result[i].Name != test.expected[i].Name || result[i].Repeat != test.expected[i].Repeat {
				t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
				break
			}
		}
	}
}
