import { urlObjectKeys } from "next/dist/shared/lib/utils"

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

export async function score(frente, deckName, reset) {
    let url = `http://localhost:3030/decks/${deckName}/cards/${frente}/score`

    if (reset) {
        url += '?reset=true'
    }

    const res = await fetch(url, {
        method: "PATCH",
    })

    if (!res.ok) {
        throw new Error("Falha ao criar a carta")
    }
}

export async function deleteCard(front, deckName) {
    const res = await fetch(`http://localhost:3030/decks/${deckName}/cards/${front}`, {
        method: "DELETE",
    }).catch(error => {
        alert(error)
    })
}