export async function list() {
    const res = await fetch('http://localhost:3030/decks')
    const decks = await res.json()

    return decks
}