import React from 'react'
import { Row, Col, Input, Select, Radio, Button, Tooltip } from 'antd'
import {
  SearchOutlined,
  PlusOutlined,
  AppstoreOutlined,
  UnorderedListOutlined,
  BorderOuterOutlined
} from '@ant-design/icons'
import { useTranslation } from 'react-i18next'

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
    <Row gutter={[16, 16]} align="middle" justify="space-between">
      <Col xs={24} sm={12} md={8} lg={6}>
        <Search
          placeholder={t('components.input.search.placeholder')}
          onChange={e => setSearchQuery(e.target.value)}
          style={{ width: '100%' }}
          prefix={<SearchOutlined />}
        />
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
                  />
                )}
                {type.label}
              </div>
            )
          }))}
          placeholder={t('components.select.room_type.placeholder')}
        />
      </Col>
      <Col xs={24} sm={12} md={8} lg={6}>
        <Radio.Group
          value={viewMode}
          onChange={e => setViewMode(e.target.value)}
          buttonStyle="solid"
          style={{ width: '100%' }}
        >
          <Radio.Button value="plan">
            <Tooltip title={t('components.tooltip.plan_view')}>
              <BorderOuterOutlined /> Plan
            </Tooltip>
          </Radio.Button>
          <Radio.Button value="grid">
            <Tooltip title={t('components.tooltip.grid_view')}>
              <AppstoreOutlined /> Grid
            </Tooltip>
          </Radio.Button>
          <Radio.Button value="list">
            <Tooltip title={t('components.tooltip.list_view')}>
              <UnorderedListOutlined /> List
            </Tooltip>
          </Radio.Button>
        </Radio.Group>
      </Col>
      <Col xs={24} sm={12} md={8} lg={6}>
        <Button
          type="primary"
          icon={<PlusOutlined />}
          onClick={() => showModal('addRoom')}
          block
        >
          {t('components.button.add_room')}
        </Button>
      </Col>
    </Row>
  )
}

export default InventoryControls
