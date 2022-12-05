package mypackage

import "fmt"

type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

func PrintMessage(text string) {
	fmt.Println("Mensaje publico")
}

func printMessagel(text string) {
	fmt.Println("Mensaje publico")
}

//https://teams.microsoft.com/l/meetup-join/19%3ameeting_YzkzMDgyODUtZmE5Ny00YjYzLWIxNTYtZTU5NDZlOTNmOWQ1%40thread.v2/0?context=%7b%22Tid%22%3a%22bf208dcb-97e8-4d43-bd72-323680bef25c%22%2c%22Oid%22%3a%22f8dc5a02-9a64-4e03-a4bd-445fa271eb43%22%7d
