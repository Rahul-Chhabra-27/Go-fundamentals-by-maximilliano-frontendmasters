package main

// ** global variable..
const url = "https://frontendmasters.com";

 // ! we can not use this way of creating variables in global...
// greet := "messgae";

func main() {
	print("Hello from GO!!\n");print("Hello from Rahul!!")

	// ! Will give an error, not a way of declaring vaiables.
	// var high;

	// * * function-scoped-variables
	// variables...
	var firstName, lastName string = "Rahul", "Chhabra";
	print(firstName," ",lastName,"\n");

	// variables....
	var nameofYourSchool = "St andrews";
	var numberOfTeachers = 20
	print(nameofYourSchool,numberOfTeachers,"\n");

	// variables....
	nameofYourCollege := "Sharda University";
	numberOfStdents := 10
	print(nameofYourCollege,numberOfStdents,"\n");

	// * * Block-scoped-variables....
	{
		hello := "hello ";
		hello += "bae";

		if isTrue := true; isTrue {
			print("Bravo!\n");
		}
	}

	greet();
	print(url,"\n");
}