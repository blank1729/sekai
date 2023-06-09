import { Button } from "@mui/base";
import { useState } from "react";
import {
  StyledButton,
  StyledInput,
  StyledModal,
} from "../components/StyledButton";

const Mui = () => {
  const [show, setShow] = useState(false);

  return (
    <main>
      <section className="container">
        <Button onClick={() => setShow(!show)}>Show modal</Button>
        <StyledModal open={show}>
          <h1>shit</h1>
        </StyledModal>
      </section>
      <StyledButton>Button</StyledButton>
      <StyledInput />
    </main>
  );
};

export default Mui;
