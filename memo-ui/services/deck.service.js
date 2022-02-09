export async function list() {
    const res = await fetch('http://localhost:3030/decks')
    const decks = await res.json()

    return decks
}

export async function create(name, description) {
    const data = {
        name,
        description
    }
    const res = await fetch('http://localhost:3030/decks', {
        method: "POST",
        body: JSON.stringify(data)
    }).catch(error => {
        alert(error)
    })
}