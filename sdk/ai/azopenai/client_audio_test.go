//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azopenai_test

import (
	"context"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/stretchr/testify/require"
)

func TestClient_GetAudioTranscription_AzureOpenAI(t *testing.T) {
	if recording.GetRecordMode() != recording.LiveMode {
		t.Skipf("Recording needs to be revisited for multipart: https://github.com/Azure/azure-sdk-for-go/issues/21598")
	}

	client := newTestClient(t, azureWhisper, withForgivingRetryOption())
	runTranscriptionTests(t, client, azureWhisperModel)
}

func TestClient_GetAudioTranscription_OpenAI(t *testing.T) {
	if recording.GetRecordMode() != recording.LiveMode {
		t.Skipf("Recording needs to be revisited for multipart: https://github.com/Azure/azure-sdk-for-go/issues/21598")
	}

	client := newOpenAIClientForTest(t)

	mp3Bytes, err := os.ReadFile(`testdata/sampledata_audiofiles_myVoiceIsMyPassportVerifyMe01.mp3`)
	require.NoError(t, err)

	args := newTranscriptionOptions(azopenai.AudioTranscriptionFormatVerboseJSON, openAIWhisperModel, mp3Bytes)
	transcriptResp, err := client.GetAudioTranscription(context.Background(), args, nil)
	require.NoError(t, err)
	require.NotEmpty(t, transcriptResp)

	require.NotEmpty(t, *transcriptResp.Text)
	require.Greater(t, *transcriptResp.Duration, float32(0.0))
	require.NotEmpty(t, *transcriptResp.Language)
	require.NotEmpty(t, transcriptResp.Segments)
	require.NotEmpty(t, transcriptResp.Segments[0])
	require.NotEmpty(t, transcriptResp.Task)
}

func TestClient_GetAudioTranslation_AzureOpenAI(t *testing.T) {
	if recording.GetRecordMode() != recording.LiveMode {
		t.Skipf("Recording needs to be revisited for multipart: https://github.com/Azure/azure-sdk-for-go/issues/21598")
	}

	client := newTestClient(t, azureWhisper, withForgivingRetryOption())
	runTranslationTests(t, client, azureWhisperModel)
}

func TestClient_GetAudioTranslation_OpenAI(t *testing.T) {
	if recording.GetRecordMode() != recording.LiveMode {
		t.Skipf("Recording needs to be revisited for multipart: https://github.com/Azure/azure-sdk-for-go/issues/21598")
	}

	client := newOpenAIClientForTest(t)

	mp3Bytes, err := os.ReadFile(`testdata/sampledata_audiofiles_myVoiceIsMyPassportVerifyMe01.mp3`)
	require.NoError(t, err)

	args := newTranslationOptions(azopenai.AudioTranslationFormatVerboseJSON, openAIWhisperModel, mp3Bytes)
	transcriptResp, err := client.GetAudioTranslation(context.Background(), args, nil)
	require.NoError(t, err)
	require.NotEmpty(t, transcriptResp)

	require.NotEmpty(t, *transcriptResp.Text)
	require.Greater(t, *transcriptResp.Duration, float32(0.0))
	require.NotEmpty(t, *transcriptResp.Language)
	require.NotEmpty(t, transcriptResp.Segments)
	require.NotEmpty(t, transcriptResp.Segments[0])
	require.NotEmpty(t, transcriptResp.Task)
}

func runTranscriptionTests(t *testing.T, client *azopenai.Client, model string) {
	mp3Bytes, err := os.ReadFile(`testdata/sampledata_audiofiles_myVoiceIsMyPassportVerifyMe01.mp3`)
	require.NoError(t, err)

	t.Run(string(azopenai.AudioTranscriptionFormatText), func(t *testing.T) {
		args := newTranscriptionOptions(azopenai.AudioTranscriptionFormatText, model, mp3Bytes)
		transcriptResp, err := client.GetAudioTranscription(context.Background(), args, nil)
		require.NoError(t, err)
		require.NotEmpty(t, transcriptResp)

		require.NotEmpty(t, *transcriptResp.Text)
		requireEmptyAudioTranscription(t, transcriptResp.AudioTranscription)
	})

	t.Run(string(azopenai.AudioTranscriptionFormatSrt), func(t *testing.T) {
		args := newTranscriptionOptions(azopenai.AudioTranscriptionFormatSrt, model, mp3Bytes)
		transcriptResp, err := client.GetAudioTranscription(context.Background(), args, nil)
		require.NoError(t, err)
		require.NotEmpty(t, transcriptResp)

		require.NotEmpty(t, *transcriptResp.Text)
		requireEmptyAudioTranscription(t, transcriptResp.AudioTranscription)
	})

	t.Run(string(azopenai.AudioTranscriptionFormatVtt), func(t *testing.T) {
		args := newTranscriptionOptions(azopenai.AudioTranscriptionFormatVtt, model, mp3Bytes)
		transcriptResp, err := client.GetAudioTranscription(context.Background(), args, nil)
		require.NoError(t, err)
		require.NotEmpty(t, transcriptResp)

		require.NotEmpty(t, *transcriptResp.Text)
		requireEmptyAudioTranscription(t, transcriptResp.AudioTranscription)
	})

	t.Run(string(azopenai.AudioTranscriptionFormatVerboseJSON), func(t *testing.T) {
		args := newTranscriptionOptions(azopenai.AudioTranscriptionFormatVerboseJSON, model, mp3Bytes)
		transcriptResp, err := client.GetAudioTranscription(context.Background(), args, nil)
		require.NoError(t, err)
		require.NotEmpty(t, transcriptResp)

		require.NotEmpty(t, *transcriptResp.Text)
		require.Greater(t, *transcriptResp.Duration, float32(0.0))
		require.NotEmpty(t, *transcriptResp.Language)
		require.NotEmpty(t, transcriptResp.Segments)
		require.NotEmpty(t, transcriptResp.Segments[0])
		require.NotEmpty(t, transcriptResp.Task)
	})

	t.Run(string(azopenai.AudioTranscriptionFormatJSON), func(t *testing.T) {
		args := newTranscriptionOptions(azopenai.AudioTranscriptionFormatJSON, model, mp3Bytes)
		transcriptResp, err := client.GetAudioTranscription(context.Background(), args, nil)
		require.NoError(t, err)
		require.NotEmpty(t, transcriptResp)

		require.NotEmpty(t, *transcriptResp.Text)
		requireEmptyAudioTranscription(t, transcriptResp.AudioTranscription)
	})
}

func runTranslationTests(t *testing.T, client *azopenai.Client, model string) {
	mp3Bytes, err := os.ReadFile(`testdata/sampledata_audiofiles_myVoiceIsMyPassportVerifyMe01.mp3`)
	require.NoError(t, err)

	t.Run(string(azopenai.AudioTranscriptionFormatText), func(t *testing.T) {
		args := newTranslationOptions(azopenai.AudioTranslationFormatText, model, mp3Bytes)
		transcriptResp, err := client.GetAudioTranslation(context.Background(), args, nil)
		require.NoError(t, err)
		require.NotEmpty(t, transcriptResp)

		require.NotEmpty(t, *transcriptResp.Text)
		requireEmptyAudioTranscription(t, transcriptResp.AudioTranscription)
	})

	t.Run(string(azopenai.AudioTranscriptionFormatSrt), func(t *testing.T) {
		args := newTranslationOptions(azopenai.AudioTranslationFormatSrt, model, mp3Bytes)
		transcriptResp, err := client.GetAudioTranslation(context.Background(), args, nil)
		require.NoError(t, err)
		require.NotEmpty(t, transcriptResp)

		require.NotEmpty(t, *transcriptResp.Text)
		requireEmptyAudioTranscription(t, transcriptResp.AudioTranscription)
	})

	t.Run(string(azopenai.AudioTranscriptionFormatVtt), func(t *testing.T) {
		args := newTranslationOptions(azopenai.AudioTranslationFormatVtt, model, mp3Bytes)
		transcriptResp, err := client.GetAudioTranslation(context.Background(), args, nil)
		require.NoError(t, err)
		require.NotEmpty(t, transcriptResp)

		require.NotEmpty(t, *transcriptResp.Text)
		requireEmptyAudioTranscription(t, transcriptResp.AudioTranscription)
	})

	t.Run(string(azopenai.AudioTranscriptionFormatVerboseJSON), func(t *testing.T) {
		args := newTranslationOptions(azopenai.AudioTranslationFormatVerboseJSON, model, mp3Bytes)
		transcriptResp, err := client.GetAudioTranslation(context.Background(), args, nil)
		require.NoError(t, err)
		require.NotEmpty(t, transcriptResp)

		require.NotEmpty(t, *transcriptResp.Text)
		require.Greater(t, *transcriptResp.Duration, float32(0.0))
		require.NotEmpty(t, *transcriptResp.Language)
		require.NotEmpty(t, transcriptResp.Segments)
		require.NotEmpty(t, transcriptResp.Segments[0])
		require.NotEmpty(t, transcriptResp.Task)
	})

	t.Run(string(azopenai.AudioTranscriptionFormatJSON), func(t *testing.T) {
		args := newTranslationOptions(azopenai.AudioTranslationFormatJSON, model, mp3Bytes)
		transcriptResp, err := client.GetAudioTranslation(context.Background(), args, nil)
		require.NoError(t, err)
		require.NotEmpty(t, transcriptResp)

		require.NotEmpty(t, *transcriptResp.Text)
		requireEmptyAudioTranscription(t, transcriptResp.AudioTranscription)
	})
}

func newTranscriptionOptions(format azopenai.AudioTranscriptionFormat, model string, mp3Bytes []byte) azopenai.AudioTranscriptionOptions {
	return azopenai.AudioTranscriptionOptions{
		Deployment:     model,
		File:           mp3Bytes,
		ResponseFormat: &format,
		Language:       to.Ptr("en"),
		Temperature:    to.Ptr[float32](0.0),
	}
}

func newTranslationOptions(format azopenai.AudioTranslationFormat, model string, mp3Bytes []byte) azopenai.AudioTranslationOptions {
	return azopenai.AudioTranslationOptions{
		Deployment:     model,
		File:           mp3Bytes,
		ResponseFormat: &format,
		Temperature:    to.Ptr[float32](0.0),
	}
}

// requireEmptyAudioTranscription checks that all the attributes are empty (aside
// from Text)
func requireEmptyAudioTranscription(t *testing.T, at azopenai.AudioTranscription) {
	// Text is always filled out for

	require.Empty(t, at.Duration)
	require.Empty(t, at.Language)
	require.Empty(t, at.Segments)
	require.Empty(t, at.Task)
}
