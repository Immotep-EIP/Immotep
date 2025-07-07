import React from 'react'
import { useTranslation } from 'react-i18next'

import { Row, Col, Input, Select, Radio, Tooltip } from 'antd'
import {
  SearchOutlined,
  PlusOutlined,
  AppstoreOutlined,
  UnorderedListOutlined,
  BorderOuterOutlined
} from '@ant-design/icons'

import { Button } from '@/components/common'

import { InventoryControlsProps } from '@/interfaces/Property/Inventory/Inventory'

const { Search } = Input

const InventoryControls: React.FC<InventoryControlsProps> = ({
  setSearchQuery,
  selectedRoomType,
  setSelectedRoomType,
  viewMode,
  setViewMode,
  showModal,
  roomTypes
}) => {
  const { t } = useTranslation()

  return (
    <section aria-labelledby="inventory-controls-title">
      <h2 id="inventory-controls-title" className="sr-only">
        {t('components.inventory.controls.search')}
      </h2>

      <Row
        gutter={[16, 16]}
        align="middle"
        justify="space-between"
        role="toolbar"
        aria-label={t('components.inventory.controls.search')}
      >
        <Col xs={24} sm={12} md={8} lg={6}>
          <div role="search">
            <Search
              placeholder={t('components.input.search.placeholder')}
              onChange={e => setSearchQuery(e.target.value)}
              style={{ width: '100%' }}
              prefix={<SearchOutlined aria-hidden="true" />}
              aria-label={t('components.inventory.controls.search')}
              id="inventory-search"
            />
          </div>
        </Col>
        <Col xs={24} sm={12} md={8} lg={6}>
          <Select
            style={{ width: '100%' }}
            value={selectedRoomType}
            onChange={setSelectedRoomType}
            options={roomTypes.map(type => ({
              ...type,
              label: (
                <div
                  style={{ display: 'flex', alignItems: 'center', gap: '8px' }}
                >
                  {type.color && (
                    <div
                      style={{
                        width: '12px',
                        height: '12px',
                        borderRadius: '50%',
                        backgroundColor: type.color
                      }}
                      aria-hidden="true"
                    />
                  )}
                  {type.label}
                </div>
              )
            }))}
            placeholder={t('components.select.room_type.placeholder')}
            aria-label={t('components.inventory.controls.filter')}
            id="room-type-filter"
          />
        </Col>
        <Col xs={24} sm={12} md={8} lg={6}>
          <fieldset
            style={{
              border: 'none',
              padding: '0',
              margin: '0'
            }}
          >
            <legend className="sr-only">
              {t('components.inventory.controls.view_mode')}
            </legend>
            <Radio.Group
              value={viewMode}
              onChange={e => setViewMode(e.target.value)}
              buttonStyle="solid"
              style={{ width: '100%' }}
              aria-label={t('components.inventory.controls.view_mode')}
            >
              <Radio.Button
                value="plan"
                aria-label={t('components.tooltip.plan_view')}
              >
                <Tooltip title={t('components.tooltip.plan_view')}>
                  <BorderOuterOutlined aria-hidden="true" />{' '}
                  {t('components.inventory.view_mode.plan')}
                </Tooltip>
              </Radio.Button>
              <Radio.Button
                value="grid"
                aria-label={t('components.tooltip.grid_view')}
              >
                <Tooltip title={t('components.tooltip.grid_view')}>
                  <AppstoreOutlined aria-hidden="true" />{' '}
                  {t('components.inventory.view_mode.grid')}
                </Tooltip>
              </Radio.Button>
              <Radio.Button
                value="list"
                aria-label={t('components.tooltip.list_view')}
              >
                <Tooltip title={t('components.tooltip.list_view')}>
                  <UnorderedListOutlined aria-hidden="true" />{' '}
                  {t('components.inventory.view_mode.list')}
                </Tooltip>
              </Radio.Button>
            </Radio.Group>
          </fieldset>
        </Col>
        <Col xs={24} sm={12} md={8} lg={6}>
          <Button
            icon={<PlusOutlined aria-hidden="true" />}
            onClick={() => showModal('addRoom')}
            block
            aria-label={t('components.inventory.controls.add_room')}
            id="add-room-button"
          >
            {t('components.button.add_room')}
          </Button>
        </Col>
      </Row>
    </section>
  )
}

export default InventoryControls
