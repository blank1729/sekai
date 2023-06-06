import { Link } from "react-router-dom";
import "./header.scss";

const Header = () => {
  return (
    <header>
      <nav className="container">
        <Link to={"/"}>
          <h1>Sekai</h1>
        </Link>
        <div className="links">
          <Link to={"/"}>Home</Link>
          <Link to={"/cart"}>Cart</Link>
        </div>
        {/* <StyledMenu theme={"dark"}>
          <MenuItem>Item 1</MenuItem>
          <MenuItem>Item 2</MenuItem>
          <MenuItem>Item 3</MenuItem>
        </StyledMenu> */}
      </nav>
    </header>
  );
};

export default Header;
