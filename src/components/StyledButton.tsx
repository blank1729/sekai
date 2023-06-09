import { Button, Input, Modal } from "@mui/base";
import styled from "styled-components";

const StyledModal = styled(Modal)`
  position: fixed;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  background-color: tomato;
  width: 40%;
  height: 25%;
  display: flex;
  justify-content: center;
  align-items: center;
  color: white;
  font-size: 1.5rem;
  border-radius: 15px;
`;

const StyledButton = styled(Button)`
  padding: 0.5rem;
  font-size: 1.25rem;
  border-radius: 0.5rem;
  background-color: #a0e7e5;
  transition: all 200ms cubic-bezier(0.15, 1.12, 0.5, 1.6);
  border: none;
  box-shadow: none;
  &:hover {
    box-shadow: 0.25rem 0.25rem black;
    transform: translate(-0.25rem, -0.25rem);
  }
  &:active {
    transform: translate(0rem, 0rem);
    box-shadow: none;
  }
`;

const StyledInput = styled(Input)`
  padding: 0.5rem;
  font-size: 1.25rem;
  border-radius: 0.5rem;
  border: none;
  outline: none;
  font-size: 2rem;
  &:focus {
    background-color: #a0e7e5;
  }
`;

export { StyledButton, StyledInput, StyledModal };
