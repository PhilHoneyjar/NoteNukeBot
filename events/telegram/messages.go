package telegram

const msgHelp = `I can save links you interested to read later. Just send them to me. 
If you send /rnd command I will send you random link and delete it from your list.
If you don't want link to be deleted just send it back to me`

const msgHello = "Howdy! \n\n" + msgHelp

const (
	msgUnknownCommand = "I don't know this command"
	msgNoSavedPages   = "Your list of links is empty right now. Save some links first"
	msgSaved          = "Your link was saved"
	msgAlredyExists   = "This link in your list now"
)
