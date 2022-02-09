import { Flex, Box, Button } from "rebass"
import { Textarea, Input, Label } from "@rebass/forms"
import MemoHeading from "../components/memoheading/memoheading"
import theme from "../theme"
import Link from "next/link"
import { useState } from "react"
import { create } from "../services/deck.service"

function CreateDeck() {
    const [name, setName] = useState("")
    const [textArea, setTextArea] = useState("")

    const onInputChange = (e) => {
        setName(e.target.value)
    }

    const onTextAreaChange = (e) => {
        setTextArea(e.target.value)
    }

    const post = () => {
        create(name, textArea)
    }
    return (
        <>
            <MemoHeading title='Criar Deck' />

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
                            htmlFor='name'>Nome:</Label>
                        <Input
                            id='name'
                            name='name'
                            value={name}
                            onChange={onInputChange}
                        />
                    </Box>
                    <Box justifyContent='center' px={2}>
                        <Label
                            fontSize={5}
                            fontWeight='bold'
                            textAlign='center'
                            htmlFor='description'>Descrição:</Label>
                        <Textarea
                            id='description'
                            name='description'
                            value={textArea}
                            onChange={onTextAreaChange}
                        />
                    </Box>
                </Flex>
                <Flex justifyContent='center' mx={-2} mb={3}>
                    <Box px={2}>
                        <Link href='/'>
                            <Button
                                style={{cursor: 'pointer'}}
                                backgroundColor={theme.darkgrey}
                                fontFamily='Roboto'>
                                Voltar
                            </Button>
                        </Link>
                    </Box>
                    <Box px={2}>
                        <Button
                            onClick={post}
                            style={{cursor: 'pointer'}}
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

export default CreateDeck