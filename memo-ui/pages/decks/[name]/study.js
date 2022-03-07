import { useState } from "react"
import { Flex, Text, Button } from "rebass"
import MemoHeading from "../../../components/memoheading/memoheading"
import { listStudy, score } from '../../../services/card.service'
import theme from "../../../theme"
import { useRouter } from "next/router"

function Study({ cards: cardsProps, name }) {
    const [cards, setCards] = useState(cardsProps);
    const [cardTurn, setCardTurn] = useState(0);
    const [showAnswer, setShowAnswer] = useState(false);
    const router = useRouter();

    const onVerRespostaClick = () => {
        setShowAnswer(true)
    }

    const onAcerteiClick = async () => {
        //Fazer request na rota PATCH /decks/{deckName}/cards/{front}/score
        await score(cards[cardTurn].Front, name)
        nextCard()
    }

    const onErreiClick = async () => {
        //Fazer request na rota PATCH /decks/{deckName}/cards/{front}/score
        //Passar queryparam ?reset=true
        await score(cards[cardTurn].Front, name, true)
        nextCard()
    }

    const nextCard = async () => {
        //Se houver outra carta para estudar chama a funcao abaixo, se nao volta pra pagina
        //do deck
        if (cards.length < cardTurn + 1)  {
            setCardTurn(cardTurn++)
            setShowAnswer(false)
        } else {
            const cards = await listStudy(name)
            if (cards) {
                setCards(cards)
                setCardTurn(0)
                setShowAnswer(false)
            } else {
                alert("Não há mais cartas para estudar no momento.")
                router.push("/");
            }
        }
        
    }

    return (
        <>
            <MemoHeading title={name} />
            <Flex 
                color="#FFF"
                justifyContent="center"
            >
                <Flex flexDirection="column">
                    <Text 
                        fontSize={[5]}
                        fontWeight='bold'
                        textAlign='center'
                    >
                        Frente: {cards[cardTurn].Front}
                    </Text>
                    {!showAnswer && <Button
                        onClick={onVerRespostaClick}
                        style={{ cursor: 'pointer' }}
                        backgroundColor={theme.discordblue}
                        fontFamily='Roboto'>
                        Ver Resposta
                    </Button>}
                    {showAnswer && 
                    <>
                    <Text 
                        fontSize={[5]}
                        fontWeight='bold'
                        textAlign='center'
                    >
                        Verso: {cards[cardTurn].Back}
                    </Text>
                    <Button
                        onClick={onErreiClick}
                        style={{ cursor: 'pointer' }}
                        backgroundColor={theme.red}
                        fontFamily='Roboto'>
                        Errei
                    </Button>
                    <Button
                        onClick={onAcerteiClick}
                        style={{ cursor: 'pointer' }}
                        backgroundColor={theme.green}
                        fontFamily='Roboto'>
                        Acertei
                    </Button>
                    </>}
                </Flex>
            </Flex>
        </>
    )
}

export async function getServerSideProps({ query }) {
    const name = query.name
    const cards = await listStudy(name)

    return {
        props: {
            name: name,
            cards: cards || []
        }
    }
}

export default Study