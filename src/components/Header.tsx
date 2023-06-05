import { Link } from "react-router-dom";
import "./header.scss";
import { StyledMenu } from "../styled-components/StyledMenu";
import { MenuItem } from "@mui/base";

const Header = () => {
  return (
    <header>
      <nav>
        <Link to={"/"}>
          <h1>Sekai</h1>
        </Link>
        <div className="links">
          <Link to={"/"}>HOME</Link>
          <Link to={"/cart"}>Cart</Link>
        </div>
        <StyledMenu theme={"dark"}>
          <MenuItem>Item 1</MenuItem>
          <MenuItem>Item 2</MenuItem>
          <MenuItem>Item 3</MenuItem>
        </StyledMenu>
      </nav>
    </header>
  );
};

export default Header;
