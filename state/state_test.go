package state

import "testing"

func Test1(t *testing.T) {
	wanted := "[V---\\[kylling rev korn HS]\\ \\________________________/ _____________________/---Ø]"
	state := MoveTo("vest")
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}

func Test2(t *testing.T) {
	wanted := "[V---\\_______________________\\ \\[kylling rev korn HS]/ ______________________/---Ø]"
	state := MoveTo("Båt")
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}

func Test3(t *testing.T) {
	wanted := "[V---\\_______________________\\ \\______________________/ [kylling rev korn HS]/---Ø]"
	state := MoveTo("Øst")
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}
