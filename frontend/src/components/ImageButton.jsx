import "../style.css";
function ImageButton(props) {
  const { backgroundImage, onClick, text } = props;

  const style = {
    backgroundImage: `url(${backgroundImage})`,
    backgroundSize: 'cover',
    backgroundPosition: 'center',
    width: '185px',
    height: '63px',
    cursor: 'pointer',
    fontFamily: 'Runescape',
    fontSize: '20px',
    color: 'yellow',
  };

  return <button style={style} onClick={onClick}>{text}</button>;
}

export default ImageButton;