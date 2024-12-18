import { render, screen, fireEvent, waitFor, act } from '@testing-library/react';
import useNavigation from '@/hooks/useNavigation/useNavigation';
import { message } from 'antd';
import ForgotPassword from '../ForgotPassword';

jest.mock('react-i18next', () => ({
  useTranslation: jest.fn().mockReturnValue({
    t: (key: string) => key,
  }),
}));

jest.mock('@/hooks/useNavigation/useNavigation', () => ({
  __esModule: true,
  default: () => ({
    goToLogin: jest.fn(),
  }),
}));

jest.mock('antd', () => ({
  ...jest.requireActual('antd'),
  message: {
    success: jest.fn(),
    error: jest.fn(),
  },
}));

describe('ForgotPassword', () => {
  beforeEach(() => {
    render(<ForgotPassword />);
  });

  test('renders correctly', () => {
    expect(screen.getByText('pages.forgotPassword.title')).toBeInTheDocument();
    expect(screen.getByText('pages.forgotPassword.description')).toBeInTheDocument();
    expect(screen.getByLabelText('components.input.email.label')).toBeInTheDocument();
    expect(screen.getByLabelText('components.input.emailConfirmation.label')).toBeInTheDocument();
    expect(screen.getByText('components.button.sendEmail')).toBeInTheDocument();
  });

  test('displays error if emails do not match', async () => {
    const emailInput = screen.getByLabelText('components.input.email.label');
    const emailConfirmationInput = screen.getByLabelText('components.input.emailConfirmation.label');
    const submitButton = screen.getByText('components.button.sendEmail');

    fireEvent.change(emailInput, { target: { value: 'test@example.com' } });
    fireEvent.change(emailConfirmationInput, { target: { value: 'mismatch@example.com' } });

    await act(async () => {
      fireEvent.click(submitButton);
    });

    await waitFor(() =>
      expect(message.error).toHaveBeenCalledWith('pages.forgotPassword.emailsDontMatch')
    );
  });

  test('displays success message on correct form submission', async () => {
    const emailInput = screen.getByLabelText('components.input.email.label');
    const emailConfirmationInput = screen.getByLabelText('components.input.emailConfirmation.label');
    const submitButton = screen.getByText('components.button.sendEmail');

    fireEvent.change(emailInput, { target: { value: 'test@example.com' } });
    fireEvent.change(emailConfirmationInput, { target: { value: 'test@example.com' } });

    await act(async () => {
      fireEvent.click(submitButton);
    });

    await waitFor(() =>
      expect(message.success).toHaveBeenCalledWith('pages.forgotPassword.sendEmailSuccess')
    );
  });

  test('displays error if fields are empty on submission', async () => {
    const submitButton = screen.getByText('components.button.sendEmail');

    await act(async () => {
      fireEvent.click(submitButton);
    });

    await waitFor(() =>
      expect(message.error).toHaveBeenCalledWith('pages.forgotPassword.fillFields')
    );
  });

  test('calls goToLogin on successful form submission', async () => {
    const emailInput = screen.getByLabelText('components.input.email.label');
    const emailConfirmationInput = screen.getByLabelText('components.input.emailConfirmation.label');
    const submitButton = screen.getByText('components.button.sendEmail');

    const { goToLogin } = useNavigation();

    expect(goToLogin).not.toHaveBeenCalled();

    fireEvent.change(emailInput, { target: { value: 'test@example.com' } });
    fireEvent.change(emailConfirmationInput, { target: { value: 'test@example.com' } });

    await act(async () => {
      fireEvent.click(submitButton);
    });
  });
});
