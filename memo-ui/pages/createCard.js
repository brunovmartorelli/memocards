import { Flex, Box, Button } from "rebass"
import { Textarea, Input, Label } from "@rebass/forms"
import MemoHeading from "../components/memoheading/memoheading"
import theme from "../theme"
import Link from "next/link"

function CreateCard() {
    return  ( 
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
        </Box>
        <Box justifyContent='center' px={2}>
            <Label
                fontSize={5}
                fontWeight='bold'
                textAlign='center'
                htmlFor='description'>Verso:</Label>
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

export default CreateCard