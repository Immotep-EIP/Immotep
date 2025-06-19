import React from 'react'

import { Skeleton } from 'antd'

import { Badge } from '@/components/common'
import style from '@/components/features/RealProperty/PropertyCard.module.css'

const CardPropertyLoader: React.FC<{ cards: number }> = ({ cards }) =>
  Array(cards)
    .fill(0)
    .map((_, index) => (
      // eslint-disable-next-line react/no-array-index-key
      <div className={style.card} key={index}>
        <div className={style.cardContentContainer}>
          <Badge.Ribbon
            text={<Skeleton.Input size="small" active />}
            color="lightgray"
          >
            <div className={style.cardPictureContainer}>
              <Skeleton.Image active className={style.cardPicture} />
            </div>
          </Badge.Ribbon>
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
    ))

export default CardPropertyLoader
