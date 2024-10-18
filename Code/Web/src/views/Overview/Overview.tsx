import React, { useEffect, useState } from 'react'

import getUsers from '@/services/api/User/GetUsers'
import style from './Overview.module.css'

const Overview: React.FC = () => {
  const [data, setData] = useState<any>(null)

  // API EXAMPLE
  useEffect(() => {
    const getInfo = async () => {
      try {
        const users = await getUsers()
        setData(users)
      } catch (error) {
        console.error('Error fetching data:', error)
      }
    }
    getInfo()
  }, [])

  return (
    <div className={style.layoutContainer}>
      {data && Array.isArray(data) ? (
        data.map((user: any) => (
          <div key={user.id}>
            <p>First Name: {user.firstname}</p>
            <p>Last Name: {user.lastname}</p>
            <p>Email: {user.email}</p>
            <hr />
          </div>
        ))
      ) : (
        <p>Loading...</p>
      )}
    </div>
  )
}

export default Overview
