import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'
import { Button, Modal, Form, Input, Upload, message } from 'antd'
import { UploadOutlined } from '@ant-design/icons'
import style from './4DocumentsTab.module.css'

const documents = [
  {
    id: 1,
    name: 'Document 1',
    date: '01/01/2021',
    preview: 'https://via.placeholder.com/150'
  },
  {
    id: 2,
    name: 'Document 2',
    date: '02/01/2021',
    preview: 'https://via.placeholder.com/150'
  },
  {
    id: 3,
    name: 'Document 3',
    date: '03/01/2021',
    preview: 'https://via.placeholder.com/150'
  },
  {
    id: 4,
    name: 'Document 4',
    date: '04/01/2021',
    preview: 'https://via.placeholder.com/150'
  },
  {
    id: 5,
    name: 'Document 5',
    date: '05/01/2021',
    preview: 'https://via.placeholder.com/150'
  },
  {
    id: 6,
    name: 'Document 6',
    date: '06/01/2021',
    preview: 'https://via.placeholder.com/150'
  },
  {
    id: 7,
    name: 'Document 7',
    date: '07/01/2021',
    preview: 'https://via.placeholder.com/150'
  },
  {
    id: 8,
    name: 'Document 8',
    date: '08/01/2021',
    preview: 'https://via.placeholder.com/150'
  },
  {
    id: 9,
    name: 'Document 9',
    date: '09/01/2021',
    preview: 'https://via.placeholder.com/150'
  },
  {
    id: 10,
    name: 'Document 10',
    date: '10/01/2021',
    preview: 'https://via.placeholder.com/150'
  },
  {
    id: 11,
    name: 'Document 11',
    date: '11/01/2021',
    preview: 'https://via.placeholder.com/150'
  },
  {
    id: 12,
    name: 'Document 12',
    date: '12/01/2021',
    preview: 'https://via.placeholder.com/150'
  },
  {
    id: 13,
    name: 'Document 13',
    date: '13/01/2021',
    preview: 'https://via.placeholder.com/150'
  },
  {
    id: 14,
    name: 'Document 14',
    date: '14/01/2021',
    preview: 'https://via.placeholder.com/150'
  },
  {
    id: 15,
    name: 'Document 15',
    date: '15/01/2021',
    preview: 'https://via.placeholder.com/150'
  },
  {
    id: 16,
    name: 'Document 16',
    date: '16/01/2021',
    preview: 'https://via.placeholder.com/150'
  },
  {
    id: 17,
    name: 'Document 17',
    date: '17/01/2021',
    preview: 'https://via.placeholder.com/150'
  },
  {
    id: 18,
    name: 'Document 18',
    date: '18/01/2021',
    preview: 'https://via.placeholder.com/150'
  },
  {
    id: 19,
    name: 'Document 19',
    date: '19/01/2021',
    preview: 'https://via.placeholder.com/150'
  },
  {
    id: 20,
    name: 'Document 20',
    date: '20/01/2021',
    preview: 'https://via.placeholder.com/150'
  },
  {
    id: 21,
    name: 'Document 21',
    date: '21/01/2021',
    preview: 'https://via.placeholder.com/150'
  },
  {
    id: 22,
    name: 'Document 22',
    date: '22/01/2021',
    preview: 'https://via.placeholder.com/150'
  }
]

const DocumentsTab: React.FC = () => {
  const { t } = useTranslation()

  const [isModalOpen, setIsModalOpen] = useState(false)

  const [form] = Form.useForm()

  const showModal = () => {
    setIsModalOpen(true)
  }

  const handleOk = () => {
    form
      .validateFields()
      .then(values => {
        console.log('Form values:', values)
        message.success(t('components.documents.success_add'))
        form.resetFields()
        setIsModalOpen(false)
      })
      .catch(errorInfo => {
        console.error('Validation Failed:', errorInfo)
      })
  }

  const handleCancel = () => {
    form.resetFields()
    setIsModalOpen(false)
  }

  return (
    <div className={style.tabContent}>
      <div className={style.buttonAddContainer}>
        <Button type="primary" onClick={showModal}>
          {t('components.button.add_document')}
        </Button>
      </div>
      <Modal
        title={t('pages.real_property_details.tabs.documents.modal_title')}
        open={isModalOpen}
        onOk={handleOk}
        onCancel={handleCancel}
      >
        <Form
          form={form}
          layout="vertical"
          name="add_document_form"
          initialValues={{ remember: true }}
        >
          <Form.Item
            label={t('components.input.document_name.label')}
            name="documentName"
            rules={[
              {
                required: true,
                message: t('components.input.document_name.error')
              }
            ]}
          >
            <Input
              placeholder={t('components.input.document_name.placeholder')}
              aria-label={t('components.input.document_name.placeholder')}
            />
          </Form.Item>

          <Form.Item
            label={t('components.input.document.label')}
            name="documentFile"
            valuePropName="fileList"
            getValueFromEvent={e => (Array.isArray(e) ? e : e?.fileList)}
            rules={[
              { required: true, message: t('components.input.document.error') }
            ]}
          >
            <Upload name="file" listType="text" beforeUpload={() => false}>
              <Button icon={<UploadOutlined />}>
                {t('components.input.document.placeholder')}
              </Button>
            </Upload>
          </Form.Item>
        </Form>
      </Modal>
      <div className={style.documentsContainer}>
        {documents.map(document => (
          <div key={document.id} className={style.documentContainer}>
            <div className={style.documentDateContainer}>
              <span>{document.date}</span>
            </div>
            <div className={style.documentPreviewContainer}>
              <img
                src={document.preview}
                alt={document.name}
                className={style.documentPreview}
              />
            </div>
            <div className={style.documentName}>
              <span>{document.name}</span>
            </div>
          </div>
        ))}
      </div>
    </div>
  )
}

export default DocumentsTab
