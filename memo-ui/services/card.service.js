export async function list(deckName) {
    const res = await fetch(`http://localhost:3030/decks/${deckName}/cards`)
    const cards = await res.json()

    return cards
}

export async function listStudy(deckName) {
    const res = await fetch(`http://localhost:3030/decks/${deckName}/study`)
    const cards = await res.json()

    return cards
}

export async function create(frente, verso, deckName) {
    const data = {
        "Front": frente,
        "Back": verso
    }
    const res = await fetch(`http://localhost:3030/decks/${deckName}/cards`, {
        method: "POST",
        body: JSON.stringify(data)
    })

    if (!res.ok) {
        throw new Error("Falha ao criar a carta")
    }
}