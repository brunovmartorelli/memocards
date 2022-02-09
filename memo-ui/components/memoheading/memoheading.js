import {Flex, Box, Heading} from "rebass"

function MemoHeading({ title, color }) {
    const defaultColor = 'white'
    return (
        <Heading
            fontFamily='Roboto'
            fontSize={[7]}
            color={color ? color : defaultColor}
        >
            <Flex justifyContent='center'>
                <Box px='auto'>
                    {title}
                </Box>
            </Flex>
            <hr/>
        </Heading>
    )
}

export default MemoHeading