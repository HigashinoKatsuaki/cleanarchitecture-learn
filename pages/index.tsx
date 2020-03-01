import { FC, useState, useEffect, Dispatch, SetStateAction } from 'react'
import axios from 'axios'

const requesTodos = async(seter: Dispatch<SetStateAction<{[key: string]: string}[]>>) => {
  const res = await axios.get<{[key: string]: string}[]>('http://localhost:9001/todos')
  console.log(res.data)

  seter(res.data)
}

const IndexPage: FC<null> = () => {
  const [todos, setTodos] = useState<{[key: string]: string}[]>([])

  useEffect(() => {
    if (todos.length == 0) {
      requesTodos(setTodos)
    }
  })

  const datas = todos.map(data => {
    return <p>{data.task}</p>
  })

  return <div>
    <p> {datas} </p>
  </div>
}

export default IndexPage
