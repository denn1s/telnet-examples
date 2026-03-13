package main

// handleIndex renders a list of series (hardcoded for now)
func handleIndex() string {
	html := `<html>
<head><title>Series</title></head>
<body>
  <h1>My Series</h1>
  <ul>
    <li>Breaking Bad - 5 seasons</li>
    <li>The Office - 9 seasons</li>
    <li>Dark - 3 seasons</li>
  </ul>
  <a id="myid" href="/add">Add new series</a> |
  <a href="/edit">Edit a series</a>

	<div id="ajax"></div>

	
	<style>
		h1 {
		color: red;
	}
	</style>

	<script type="application/javascript">
			const a = document.getElementById('myid')
			console.log(a)

			a.addEventListener('click', async (event) => {
				event.preventDefault()

				const response = await fetch('/add')	
				const object = await response.json()

				console.log('object', object)

				const h1 = document.createElement('h1')
				h1.append(object.title)
				h1.style.color = 'pink'
				h1.style.userSelect = 'none'
				h1.addEventListener('click', () => {
					h1.style.color = h1.style.color === 'pink' ? 'red' : 'pink'
				})

				document.getElementById('ajax').append(h1)


			})



	</script>
</body>
</html>`

	return buildResponse("666 OK", html)
}

// handleAddForm renders a form to add a new series
func handleAddForm() string {
	html := `{
		"title": "Add New Series"
	}`

	return buildResponse("200 OK", html)
}

// handleEditForm renders a form to update an existing series
func handleEditForm() string {
	html := `<html>
<head><title>Edit Series</title></head>
<body>
  <h1>Edit Series</h1>
  <form method="POST" action="/series/update">
    <input type="hidden" name="id" value="1">
    <label>Name: <input type="text" name="name" value="Breaking Bad"></label><br><br>
    <label>Seasons: <input type="number" name="seasons" value="5"></label><br><br>
    <button type="submit">Update</button>
  </form>
  <br><a href="/">Back</a>
</body>
</html>`

	return buildResponse("200 OK", html)
}
