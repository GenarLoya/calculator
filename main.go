package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Contact struct {
	firstName   string
	lastName    string
	email       string
	phoneNumber string
	state       string
	business    bool
}

var contacts []Contact

var app = tview.NewApplication()

var text = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText("(a) to add a new contact \n(q) to quit")

var form = tview.NewForm()

func addContactForm() {
	contact := Contact{}

	form.AddInputField("First Name", "", 20, nil, func(firstName string) {
		contact.firstName = firstName
	})

	form.AddInputField("Last Name", "", 20, nil, func(lastName string) {
		contact.lastName = lastName
	})

	form.AddInputField("Email", "", 20, nil, func(email string) {
		contact.email = email
	})

	form.AddInputField("Phone", "", 20, nil, func(phone string) {
		contact.phoneNumber = phone
	})

	form.AddCheckbox("Business", false, func(business bool) {
		contact.business = business
	})

	form.AddButton("Save", func() {
		contacts = append(contacts, contact)
		pages.SwitchToPage("Menu")
	})
}

var pages = tview.NewPages()

var contactsList = tview.NewList().ShowSecondaryText(false)

func addContactList() {
	contactsList.Clear()
	for index, contact := range contacts {
		contactsList.AddItem(contact.firstName+" "+contact.lastName, " ", rune(49+index), nil)
	}
}

var contactText = tview.NewTextView()

var flex = tview.NewFlex()

func main() {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			app.Stop()
		} else if event.Rune() == 97 {
			form.Clear(true)
			addContactForm()
			pages.SwitchToPage("Add Contact")
		}
		return event
	})

	pages.AddPage("Menu", text, true, true)
	pages.AddPage("Add Contact", form, true, false)

	flex.AddItem(tview.NewBox().SetBorder(true), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true), 0, 1, false)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			app.Stop()
		} else if event.Rune() == 97 {
			addContactForm()
			pages.SwitchToPage("Add Contact")
		}
		return event
	})

	contactsList.SetSelectedFunc(func(index int, name string, second_name string, shortcut rune) {
		contactText.SetText(&contacts[index])
	})

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
