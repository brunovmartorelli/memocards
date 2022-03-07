import { Flex, Box, Button } from "rebass"
import { Input, Label } from "@rebass/forms"
import MemoHeading from "../../../components/memoheading/memoheading"
import theme from "../../../theme"
import Link from "next/link"
import { useState } from "react"
import { create } from "../../../services/card.service"
import { useRouter } from "next/router"

function CreateCard() {
    const router = useRouter();
    const { name: deckName } = router.query
    const [frente, setFrente] = useState("")
    const [verso, setVerso] = useState("")
    
    const onFrenteChange = (e) => {
        setFrente(e.target.value)
    }
    const onVersoChange = (e) => {
        setVerso(e.target.value)
    }
    const post = async () => {
       try {
            await create(frente, verso, deckName);
       } catch(error) {
            alert(error)
            return
       }

       router.push(`/decks/${deckName}`);
    }
    return (
        <>
            <MemoHeading title='Criar Carta' />

            <Box
                color='#CFCFCF'
                fontFamily='Roboto'
                as='form'
                onSubmit={e => e.preventDefault()}
                py={3}>
                <Flex justifyContent='center' mx={-2} mb={3}>
                    <Box justifyContent='center' px={2}>
                        <Label
                            fontSize={5}
                            fontWeight='bold'
                            textAlign='center'
                            htmlFor='name'>Frente:</Label>

                        <Input
                            id='frente'
                            name='frente'
                            value={frente}
                            onChange={onFrenteChange}
                        />
                    </Box>
                    <Box justifyContent='center' px={2}>
                        <Label
                            fontSize={5}
                            fontWeight='bold'
                            textAlign='center'
                            htmlFor='description'>Verso:</Label>
                        <Input
                            id='verso'
                            name='verso'
                            value={verso}
                            onChange={onVersoChange}
                        />
                    </Box>
                </Flex>
                <Flex justifyContent='center' mx={-2} mb={3}>
                    <Box px={2}>
                        <Link href='/'>
                            <Button
                                style={{ cursor: 'pointer' }}
                                backgroundColor={theme.darkgrey}
                                fontFamily='Roboto'>
                                Voltar
                            </Button>
                        </Link>
                    </Box>
                    <Box px={2}>
                        <Button
                            onClick={post}
                            style={{ cursor: 'pointer' }}
                            backgroundColor={theme.discordblue}
                            fontFamily='Roboto'>
                            Criar
                        </Button>
                    </Box>
                </Flex>

            </Box>
        </>

    )
}

export default CreateCard