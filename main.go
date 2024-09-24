package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Contact struct {
	name, phone, address string
}

var listcontact []Contact

func main() {
	var contact Contact
	for {
		clearScreen()
		fmt.Println("=========================")
		fmt.Println("=      Contacs App      =")
		fmt.Println("=========================")

		fmt.Println("Menu : ")
		fmt.Println("1. Tampilkan Kontak")
		fmt.Println("2. Cari Kontak")
		fmt.Println("3. Tambahkan Kontak")
		fmt.Println("4. Edit Kontak")
		fmt.Println("5. Hapus Kontak Kontak")
		fmt.Println("6. Keluar")
		fmt.Println()
		fmt.Printf("Masukan pilihan: ")
		input := inputUser()
		clearScreen()

		switch input {
		case "1":
			getAllContact()
		case "2":
			fmt.Print("Masukan nama kontak : ")
			input := inputUser()
			getContactByName(input)
		case "3":
			fmt.Println("Masukan kontak baru")
			fmt.Printf("%-10s : ", "Nama")
			inputNama := inputUser()
			fmt.Printf("%-10s : ", "No.Hp")
			inputPhone := inputUser()
			fmt.Printf("%-10s : ", "Alamat")
			inputAddress := inputUser()

			contact.addContact(inputNama, inputPhone, inputAddress)
		case "4":
			fmt.Println("Masukan nama kontak yang ingin kamu ubah")
			fmt.Printf("%-18s : ", "Nama kontak")
			nameContact := inputUser()
			fmt.Printf("%-18s : ", "Nama kontak baru")
			newName := inputUser()
			fmt.Printf("%-18s : ", "No.Hp baru")
			newPhone := inputUser()
			fmt.Printf("%-18s : ", "Alamat baru")
			newAddress := inputUser()

			newContact := Contact{
				name:    newName,
				phone:   newPhone,
				address: newAddress,
			}

			contact.editContact(nameContact, newContact)
		case "5":
			fmt.Println("Masukan nama kontak yang akan dihapus")
			fmt.Print("nama kontak: ")
			input := inputUser()

			deleteContact(input)
		case "6":
			fmt.Println("Anda keluar....")
			return
		default:
			fmt.Println("Pilihan yang kamu masukan salah!")
		}

		fmt.Print("\nTekan enter untuk kembali ke menu utama")
		inputUser()
	}
}

// function to display contacts
func getAllContact() {
	if len(listcontact) == 0 {
		fmt.Println("Tidak ada kontak yang tersedia!")
	} else {
		fmt.Printf("%-20s %-10s %s\n", "Name", "Phone", "Address")
		for _, c := range listcontact {
			fmt.Printf("%-20s %-10s %s\n", c.name, c.phone, c.address)
		}
	}
}

// function to display contacts based on name search
func getContactByName(name string) {
	index := findIndexContact(name)
	if index < 0 {
		fmt.Printf("\nKontak dengan nama %s tidak ada!", name)
	} else {
		fmt.Printf("%-20s %-20s %s\n", "Name", "Phone", "Address")
		fmt.Printf("%-20s %-20s %s\n", listcontact[index].name, listcontact[index].phone, listcontact[index].address)
	}
}

// function to add contacts
func (c *Contact) addContact(name, phone, address string) {
	c.name = name
	c.phone = phone
	c.address = address

	listcontact = append(listcontact, *c)
	fmt.Println("\nBerhasil menambahkan kontak")
}

// function to delete contacts
func deleteContact(name string) {
	index := findIndexContact(name)
	if index < 0 {
		fmt.Println("Contact Not Found!")
	} else {
		listcontact = append(listcontact[:index], listcontact[index+1:]...)
		fmt.Println("\nBerhasil menghapus kontak")
	}
}

// function to change contacts
func (c *Contact) editContact(name string, contact Contact) {
	index := findIndexContact(name)
	if index < 0 {
		fmt.Printf("\nKontak dengan nama %s tidak ada!", name)
	} else {
		listcontact[index].name = contact.name
		listcontact[index].phone = contact.phone
		listcontact[index].address = contact.address
		fmt.Println("\nBerhasil mengubah kontak")
	}
}

// function to search for index
func findIndexContact(name string) int {
	for i, c := range listcontact {
		if c.name == name {
			return i
		}
	}
	return -1
}

// function to get input from the user
func inputUser() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return input
}

// function to clean the screen
func clearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
