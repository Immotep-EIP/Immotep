import React from 'react'
import { useLocation } from 'react-router-dom'
import { Button, Tooltip } from 'antd'
import { useTranslation } from 'react-i18next'

import returnIcon from '@/assets/icons/return.png'

import style from './RealPropertyDetails.module.css'

const RealPropertyDetails: React.FC = () => {
  const { t } = useTranslation()
  const location = useLocation()
  const { id } = location.state || {}

  return (
    <div className={style.pageContainer}>
      <span>Details for Real Property ID: {id}</span>
    </div>
  )
}

export default RealPropertyDetails
