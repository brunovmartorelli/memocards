const Card = ({front, back, score, ReviewedAt }) => {
    return (
        <div>
            <span>Front: {front}</span>
            <span>Back: {back}</span>
            
        </div>
    )
}

export default Card;