function BackgroundImage(props) {
  const { backgroundImage, children } = props;

  const style = {
    backgroundImage: `url(${backgroundImage})`,
    backgroundSize: '100% 100%',
    backgroundPosition: 'top',
    width: '100vw',
    height: '100vh',
    margin: '0',
  };

  return <div style={style}>{children}</div>;
}

export default BackgroundImage