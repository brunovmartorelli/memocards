import { list } from "./../services/deck.service"

function DeckList({decks}) {
    console.log(decks)
    return (
        <ul>
            {decks.map((deck, index) => {
              return <li key={index}>{deck.name}</li>
            })}
        </ul>
    )
}

export async function getStaticProps() {
    const decks = await list()
    return {
        props: {
            decks: decks
        }
    }
}

export default DeckList