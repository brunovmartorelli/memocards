export async function list(deckName) {
    const res = await fetch(`http://localhost:3030/decks/${deckName}/cards`)
    const cards = await res.json()

    return cards
}

export async function create(frente, verso) {
    const data = {
        frente,
        verso
    }
    const res = await fetch('http://localhost:3030/cards', {
        method: "POST",
        body: JSON.stringify(data)
    }).catch(error => {
        alert(error)
    })
}