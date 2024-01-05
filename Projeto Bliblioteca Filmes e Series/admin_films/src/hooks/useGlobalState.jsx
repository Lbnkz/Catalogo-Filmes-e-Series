import { createGlobalState } from "react-hooks-global-state";

const { useGlobalState, setGlobalState } = createGlobalState({
  activeIndex: 0,
  drawerOpen: true,
});

export { useGlobalState, setGlobalState };
