import React from 'react'
import { Skeleton } from 'antd'

import style from '@/views/RealProperty/RealProperty.module.css'

interface CardPropertyLoaderProps {
  cards: number
}

const CardPropertyLoader: React.FC<CardPropertyLoaderProps> = ({ cards }) => (
  <div className={style.cardsContainer}>
    {Array(cards)
      .fill(0)
      .map((_, index) => (
        // eslint-disable-next-line react/no-array-index-key
        <div className={style.card} key={index}>
          <div className={style.statusContainer}>
            <Skeleton.Input size="small" active />
            <Skeleton.Input size="small" active />
          </div>
          <div className={style.pictureContainer}>
            <Skeleton.Avatar
              active
              shape="circle"
              size={100}
              className={style.picture}
            />
          </div>
          <div className={style.informationsContainer}>
            <div className={style.informations}>
              <Skeleton.Input style={{ height: 20 }} active />
            </div>
            <div className={style.informations}>
              <Skeleton.Input style={{ height: 20 }} active />
            </div>
            <div className={style.informations}>
              <Skeleton.Input style={{ height: 20 }} active />
            </div>
            <div className={style.informations}>
              <Skeleton.Input style={{ height: 20 }} active />
            </div>
          </div>
        </div>
      ))}
  </div>
)

export default CardPropertyLoader
