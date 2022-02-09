import DeckList from "../components/deckList/deckList"
import { list } from "./../services/deck.service"
import { withTheme } from '@emotion/react'
import MemoHeading from "./../components/memoheading/memoheading"

const Index = ({ decks }) => {
    return (
        <>
            <MemoHeading title='MEMOCARDS'/>
            <DeckList decks={decks} />
        </>
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

export default withTheme(Index)