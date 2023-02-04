import React from 'react'





function Search() {
  return (
    <form onSubmit={(e)=>{e.preventDefault()}} className='ml-20 flex items-center shadow-black border-2 border-gray-500 rounded-3xl h-10'>
        <input placeholder='search' type="text" className='text-xl outline-none ml-4'/>
        <button type='submit' className='active:scale-95 hover:scale-105 duration-150' >
            <img src="/search.svg" className='w-8 h-8 mx-4' alt="" />
        </button>
    </form>
  )
}

export default Search