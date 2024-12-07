import React from 'react'
import { useTranslation } from 'react-i18next'
import { CloseOutlined } from '@ant-design/icons'
import { Button, Card, Form, Input, Space, Typography } from 'antd'
import addIcon from '@/assets/icons/plus.png'
import style from './3InventoryTab.module.css'

const InventoryTab: React.FC = () => {
  const { t } = useTranslation()
  const [form] = Form.useForm()

  return (
    <div className={style.tabContent}>
      <Form
        form={form}
        name="dynamic_form_complex"
        className={style.roomsContainer}
        autoComplete="off"
        initialValues={{ items: [{}] }}
        layout="vertical"
      >
        <Form.List name="items">
          {(fields, { add, remove }) => (
            <div className={style.roomsContainer}>
              {fields.map(field => (
                <Card
                  className={style.roomContainer}
                  size="small"
                  title={
                    <div
                      style={{
                        display: 'flex',
                        justifyContent: 'flex-start',
                        width: '100%'
                      }}
                    >
                      <Input
                        placeholder={t(
                          'components.input.room_name.placeholder'
                        )}
                        style={{ width: '60%' }}
                      />
                    </div>
                  }
                  key={field.key}
                  extra={
                    <CloseOutlined
                      onClick={() => {
                        remove(field.name)
                      }}
                    />
                  }
                >
                  <Form.Item label={t('pages.realPropertyDetails.tabs.inventory.list_object_name')}>
                    <Form.List name={[field.name, 'list']}>
                      {(subFields, subOpt) => (
                        <div className={style.stuffsContainer}>
                          {subFields.map(subField => (
                            <div
                              key={subField.key}
                              className={style.stuffContainer}
                            >
                              <Input
                                placeholder={t(
                                  'components.input.stuff_name.placeholder'
                                )}
                                style={{ width: '90%' }}
                              />
                              <CloseOutlined
                                onClick={() => {
                                  subOpt.remove(subField.name)
                                }}
                              />
                            </div>
                          ))}
                          <Card
                            className={style.addStuffButton}
                            onClick={() => subOpt.add()}
                          >
                            <img
                              src={addIcon}
                              alt="add"
                              style={{ width: '30px', height: '30px' }}
                            />
                          </Card>
                        </div>
                      )}
                    </Form.List>
                  </Form.Item>
                </Card>
              ))}

              <div
                className={style.roomContainer}
                onClick={() => add()}
                style={{ cursor: 'pointer' }}
                tabIndex={0}
                role="button"
                onKeyDown={e => {
                  if (e.key === 'Enter') {
                    add()
                  }
                }}
              >
                <div className={style.addRoomButton}>
                  <img
                    src={addIcon}
                    alt="add"
                    style={{ width: '50px', height: '50px' }}
                  />
                </div>
              </div>
            </div>
          )}
        </Form.List>
      </Form>
    </div>
  )
}

export default InventoryTab
