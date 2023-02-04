import React from 'react'
import { item } from '../../types/item'


function MenuItem(props: item) {
	return (
		<div className="menu-item w-[200px] rounded-lg bg-gray-100 p-3 hover:scale-105 duration-200">
              <div className="logo flex justify-center items-center">
                <img src={props.img} className='bg-gray-200 rounded-lg' alt="" />
              </div>
              <div className="grid gap-y-2 mt-4">
                <div className="flex items-center justify-between">
                  <h3 className="text-xl font-bold">{props.id}</h3>
                  <img src="/vegetarian.png" className="w-5 h-5" alt="" />
                </div>
                <div className="flex justify-between items-center">
                  <p className='text-sm'>{props.name}</p>
                  <p className="price font-semibold">&#x20B9; {props.price}</p>
                </div>
              </div>
            </div>
	)
}

export default MenuItem