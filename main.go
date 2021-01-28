package main

import "fmt"

func main() {
	fmt.Println(Pagination(1, 12345, 10))
}

//Pagination ..
func Pagination(currentPage, totalRecords, PerPage int) string {
	buttonBox := ""
	totalPages := totalRecords / PerPage
	if totalPages != 0 {
		buttonBox = `<div class=""> `
		if currentPage == 1 {
			buttonBox = buttonBox + `<button class="" disabled>Previous</button>`
		} else {
			buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-1, "Previous")
		}
		if currentPage == 1 {
			if totalPages == 0 {
				buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, currentPage)
			} else if totalPages == 1 {
				buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, currentPage)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+1, currentPage+1)
			} else if totalPages == 2 {
				buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, currentPage)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+1, currentPage+1)
			} else if totalPages == 3 {
				buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, currentPage)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+1, currentPage+1)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, totalPages, totalPages)
			} else if totalPages == 4 {
				buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, currentPage)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+1, currentPage+1)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+2, currentPage+2)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, totalPages, totalPages)
			} else if totalPages >= 5 {
				if totalPages == 5 {
					buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, currentPage)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+1, currentPage+1)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+2, currentPage+2)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+3, currentPage+3)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, totalPages, totalPages)
				} else {
					buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button></a>`, currentPage)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+1, currentPage+1)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+2, currentPage+2)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+3, currentPage+3)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+4, currentPage+4)
					buttonBox = buttonBox + fmt.Sprintf(`<button class="" disabled>%v</button></a>`, "...")
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, totalPages, totalPages)
				}
			}

		} else if currentPage == 2 {
			if totalPages == 2 {
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-1, currentPage-1)
				buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, currentPage)
			} else if totalPages == 3 {
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-1, currentPage-1)
				buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, currentPage)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, totalPages, totalPages)
			} else if totalPages == 4 {
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-1, currentPage-1)
				buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button></a>`, currentPage)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+1, currentPage+1)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, totalPages, totalPages)
			} else if totalPages >= 5 {
				if totalPages == 5 {
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-1, currentPage-1)
					buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, currentPage)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+1, currentPage+1)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+2, currentPage+2)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, totalPages, totalPages)
				} else {
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-1, currentPage-1)
					buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, currentPage)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+1, currentPage+1)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+2, currentPage+2)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+3, currentPage+3)
					buttonBox = buttonBox + fmt.Sprintf(`<button class="" disabled>%v</button></a>`, "...")
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, totalPages, totalPages)
				}
			}
		} else if currentPage == 3 {
			if totalPages == 3 {
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-2, currentPage-2)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-1, currentPage-1)
				buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, totalPages)
			} else if totalPages == 4 {
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-2, currentPage-2)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-1, currentPage-1)
				buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, currentPage)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, totalPages, totalPages)
			} else if totalPages >= 5 {
				if totalPages == 5 {
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-2, currentPage-2)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-1, currentPage-1)
					buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, currentPage)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+1, currentPage+1)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, totalPages, totalPages)
				} else {
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-2, currentPage-2)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-1, currentPage-1)
					buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, currentPage)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+1, currentPage+1)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+2, currentPage+2)
					buttonBox = buttonBox + fmt.Sprintf(`<button class="" disabled>%v</button></a>`, "...")
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, totalPages, totalPages)
				}
			}
		} else if currentPage == 4 { //Page 4
			if totalPages == 4 {
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-3, currentPage-3)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-2, currentPage-2)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-1, currentPage-1)
				buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, currentPage)
			} else if totalPages >= 5 {
				if totalPages == 5 {
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-3, currentPage-3)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-2, currentPage-2)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-1, currentPage-1)
					buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, currentPage)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, totalPages, totalPages)
				} else {
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-3, currentPage-3)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-2, currentPage-2)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-1, currentPage-1)
					buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, currentPage)
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+1, currentPage+1)
					buttonBox = buttonBox + fmt.Sprintf(`<button class="" disabled>%v</button></a>`, "...")
					buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, totalPages, totalPages)
				}
			}
		} else if currentPage == 5 {
			if totalPages == 5 {
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-4, currentPage-4)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-3, currentPage-3)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-2, currentPage-2)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-1, currentPage-1)
				buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, currentPage)
			} else {
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-4, currentPage-4)
				buttonBox = buttonBox + fmt.Sprintf(`<button class="" disabled>%v</button></a>`, "...")
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-1, currentPage-1)
				buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, currentPage)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+1, currentPage+1)
				buttonBox = buttonBox + fmt.Sprintf(`<button class="" disabled>%v</button></a>`, "...")
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, totalPages, totalPages)
			}
		} else if currentPage >= 6 {
			if totalPages == currentPage {
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, "1", "1")
				buttonBox = buttonBox + fmt.Sprintf(`<button class="" disabled>%v</button></a>`, "...")
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-2, currentPage-2)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-1, currentPage-1)
				buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, currentPage)
			} else if currentPage == (totalPages - 2) {
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, "1", "1")
				buttonBox = buttonBox + fmt.Sprintf(`<button class="" disabled>%v</button></a>`, "...")
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-1, currentPage-1)
				buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, currentPage)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+1, currentPage+1)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, totalPages, totalPages)
			} else if currentPage == (totalPages - 1) {
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, "1", "1")
				buttonBox = buttonBox + fmt.Sprintf(`<button class="" disabled>%v</button></a>`, "...")
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-1, currentPage-1)
				buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, currentPage)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, totalPages, totalPages)
			} else {
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, "1", "1")
				buttonBox = buttonBox + fmt.Sprintf(`<button class="" disabled>%v</button></a>`, "...")
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage-1, currentPage-1)
				buttonBox = buttonBox + fmt.Sprintf(`<button class="">%v</button>`, currentPage)
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+1, currentPage+1)
				buttonBox = buttonBox + fmt.Sprintf(`<button class="" disabled>%v</button></a>`, "...")
				buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, totalPages, totalPages)
			}
		}
		if currentPage == totalPages {
			buttonBox = buttonBox + `<button class="" disabled>Next</button>`
		} else {
			buttonBox = buttonBox + fmt.Sprintf(`<a href="%v" ><button class="">%v</button></a>`, currentPage+1, "Next")
		}
		buttonBox = buttonBox + `</div>`
	}

	return buttonBox
}
