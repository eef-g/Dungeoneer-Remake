function BackgroundImage(props) {
  const { backgroundImage, children } = props;

  const style = {
    backgroundImage: `url(${backgroundImage})`,
    backgroundSize: '100% 100%',
    backgroundPosition: 'top',
    width: '409px',
    height: '627px',
    margin: '0',
  };

  return <div style={style}>{children}</div>;
}

export default BackgroundImage