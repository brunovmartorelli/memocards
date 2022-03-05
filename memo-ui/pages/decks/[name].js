import MemoHeading from "../../components/memoheading/memoheading"
import { list } from '../../services/card.service'
import CardList from "../../components/cardList/cardList"

function DeckPage({cards, name}) {

    return (
        <>
            <MemoHeading title={name} />
            <CardList cards={cards} />
        </>
    )
}

export async function getServerSideProps({query}) {
    const name = query.name
    const cards = await list(name)
    return {
        props: {
            name: name,
            cards: cards
        }
    }
}

export default DeckPage