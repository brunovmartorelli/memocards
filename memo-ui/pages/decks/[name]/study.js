import { useState } from "react"
import { Flex, Text, Button } from "rebass"
import MemoHeading from "../../../components/memoheading/memoheading"
import { listStudy } from '../../../services/card.service'
import theme from "../../../theme"

function Study({ cards, name }) {
    const [cardTurn, setCardTurn] = useState(0);

    const onVerRespostaClick = () => {
        //Mostrar o back da carta e esconder o botao de ver resposta e mostrar
        //botoes de Acertei e Errei
    }

    const onAcerteiClick = () => {
        //Fazer request na rota PATCH /decks/{deckName}/cards/{front}/score
    }

    const onErreiClick = () => {
        //Fazer request na rota PATCH /decks/{deckName}/cards/{front}/score
        //Passar queryparam ?reset=true
    }

    const nextCard = () => {
        //Se houver outra carta para estudar chama a funcao abaixo, se nao volta pra pagina
        //do deck
        //setCardTurn(cardTurn++)
    }

    return (
        <>
            <MemoHeading title={name} />
            <Flex
                justifyContent="center"
            >
                <Flex flexDirection="column">
                    <Text fontSize={[5]}
                        fontWeight='bold'
                        textAlign='center'>Front: {cards[cardTurn].front}</Text>
                    <Button
                        style={{ cursor: 'pointer' }}
                        backgroundColor={theme.discordblue}
                        fontFamily='Roboto'>
                        Ver Resposta
                    </Button>
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
            cards: cards
        }
    }
}

export default Study