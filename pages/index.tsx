import { FC, useState, useEffect, Dispatch, SetStateAction } from 'react'
import axios from 'axios'
import { InjectedFormikProps, withFormik, Form, Field } from 'formik';

const requesTodos = async(seter: Dispatch<SetStateAction<{[key: string]: string}[]>>) => {
  const res = await axios.get<{[key: string]: string}[]>('http://localhost:9001/todos')
  console.log(res.data)

  seter(res.data)
}

const createTodo = async(value: FormValues) => {
  console.log(value.task)
  await axios.post('http://localhost:9001/todos', {
    status: 'in progress',
    task: value.task
  })
}

const createValidate: (values: FormValues) => object = (values) => {
  const error: {[key: string]: string} = {}
  if (!values.task) {
    error.task = 'required'
  }
  return error
}

interface FormProps {
  task?: string
}

interface FormValues {
  task?: string
}

const CreateForm: FC<InjectedFormikProps<FormProps, FormValues>> = (props) => {
  const { touched, errors, isSubmitting } = props
  return <div>
    <h3>Todo追加</h3>
    <Form>
      <p><Field type="text" name="task" /></p>
      { touched.task && errors.task && <div>{errors.task}</div> }
      <button type="submit" disabled={isSubmitting}>追加</button>
    </Form>
    </div>
}

const CreateFormWithFormik = withFormik<FormProps, FormValues>({
  mapPropsToValues() {
    return {
      task: ''
    }
  },
  validate: createValidate,
  handleSubmit: createTodo,
})(CreateForm)

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
    <p>
      <div>
        <CreateFormWithFormik />
      </div>
    </p>
  </div>
}

export default IndexPage
