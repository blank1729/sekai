import React from 'react'
import { useAppDispatch, useAppSelector } from '../app/hooks'
import { incremented, decremented, reset } from '../features/counter/counter-slice';

function Counter() {
  const count = useAppSelector(state => state.counter.value)
  const dispatch = useAppDispatch();
	function handleIncrement(){
		dispatch(incremented())
	}
    return (
    <main>
        <div>value is {count}</div>
        <button onClick={()=> dispatch(decremented())}>Decrement</button>
        <button onClick={()=> dispatch(reset())}>Reset</button>
        <button onClick={()=> dispatch(incremented())}>increment</button>
    </main>
  )
}

export default Counter