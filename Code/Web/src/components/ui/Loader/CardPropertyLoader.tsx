import React from 'react'
import { Skeleton } from 'antd'

import style from '@/views/RealProperty/RealProperty.module.css'

const CardPropertyLoader: React.FC<{ cards: number }> = ({ cards }) => (
  <div className={style.cardsContainer}>
    {Array(cards)
      .fill(0)
      .map((_, index) => (
        // eslint-disable-next-line react/no-array-index-key
        <div className={style.card} key={index}>
          <div className={style.cardContentContainer}>
            <div className={style.cardPictureContainer}>
              <Skeleton.Image active className={style.cardPicture} />
            </div>
            <div className={style.cardInfoContainer}>
              <b className={style.cardText}>
                <Skeleton.Input size="small" active />
              </b>
              <div className={style.cardAddressContainer}>
                <Skeleton.Avatar size={20} active />
                <span className={style.cardText}>
                  <Skeleton.Input size="small" active />
                </span>
              </div>
            </div>
          </div>
        </div>
      ))}
  </div>
)

export default CardPropertyLoader
