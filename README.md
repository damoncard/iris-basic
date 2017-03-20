# Todo CRUD

First meet with iris framework with Todo List CRUD.

## Scope

* Create Todo List Application
	* The **todo** page should display:
		* All list items in todo-list. Each item should have:
			* Title
			* Date to complete
			* A field that indicates if the items is done
				* Yes, if finished.
				* No, if not finished.
			* Delete button
		* Adding item section which have:
			* Title to a new list
			* Deadline for the task
			* Status of the task:
				* Yes, if it's already finished.
				* No, if it's not yet finish.
		* Counter display:
			* Show the number of the *completed* task
			* Show the number of the *uncompleted* task
	* The **edit** page should display:
		* **input text** for title of the item
		* **date** for the date of the item
		* **status** of the item
		Note: **All** field must be completed before submiting
* Create Basic Authentication	
	*Have **2** levels of authentication
		* **Admin** for CRUD function
		* **Guest** for R function

### Technologies

* [Golang v.1.8](https://github.com/golang/go)
* [Iris framework](https://github.com/kataras/iris)
* [Bootstrap](https://github.com/twbs/bootstrap)
* [BoltDB](https://github.com/boltdb/bolt)

### Setup

1. ~~Create Repo~~
2. ~~Setup project on ubuntu (VMware)~~
3. Download and Install
	1. ~~Go v.1.8~~
	2. ~~Iris framework~~
	3. Bootstrap
	4. BoltDB
4. ~~Initialize project with iris framework~~
5. Use **httprouter** adapter for routing
5. Use **view.HTML** to set template
6. Use **sessioin** for authentication