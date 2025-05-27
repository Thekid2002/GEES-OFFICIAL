import os
import torch
import numpy as np
import pandas as pd
import joblib
import datetime
from flask import Flask, request, jsonify
from ModelTrainer import GestureModel

app = Flask(__name__)

# Load model and preprocessing tools at startup
def load_model():
    global model, label_encoder, scaler, input_dim, output_dim
    
    label_encoder = joblib.load("gesture_label_encoder.pkl")
    scaler = joblib.load("scaler.pkl")
    input_dim = scaler.n_features_in_
    output_dim = len(label_encoder.classes_)

    model = GestureModel(input_dim, output_dim)
    model.load_state_dict(torch.load("gesture_model.pt"))
    model.eval()

# Call load_model when the application starts
load_model()

@app.route('/predict', methods=['POST'])
def predict():
    try:
        feature_data = request.json
        
        # Handle nested DataPoints separately if present
        if 'DataPoints' in feature_data:
            # Remove DataPoints array before creating DataFrame
            data_points = feature_data.pop('DataPoints', None)
            # Log for debugging
            print(f"Removed DataPoints from feature data, found {len(data_points) if data_points else 0} points")
        
        # Convert key names if needed
        converted_data = {}
        for key, value in feature_data.items():
            new_key = key.lower()
            converted_data[new_key] = value
        
        # Create DataFrame from the flattened data
        dataFrame = pd.DataFrame([converted_data])
        
        # Remove columns that shouldn't be part of prediction features
        if 'id' in dataFrame.columns:
            dataFrame = dataFrame.drop(columns=['id'])
        if 'gestureid' in dataFrame.columns:
            dataFrame = dataFrame.drop(columns=['gestureid'])
        if 'datapointscount' in dataFrame.columns:
            dataFrame = dataFrame.drop(columns=['datapointscount'])
        
        try:
            data_np = dataFrame.values.astype(np.float32)
            # This function scales the values, without affecting the original Relative sizes
            data_scaled = scaler.transform(data_np)
            # Convert to tensor
            input_tensor = torch.tensor(data_scaled, dtype=torch.float32)
            
        except Exception as e:
            return jsonify({"error": f"Error scaling data: {str(e)}", "details": dataFrame.to_dict()}), 400


        # Disables gradient calculation
        with torch.no_grad():
            # Make prediction
            output = model(input_tensor)
            prediction = torch.argmax(output, dim=1).item()
            predicted_gesture_id = int(label_encoder.inverse_transform([prediction])[0])
            
            # Get confidence scores
            probabilities = torch.nn.functional.softmax(output, dim=1)[0]
            confidence = float(probabilities[prediction].item())

            # Log all probabilities with labels for debugging
            all_probs = {}
            all_classes = label_encoder.classes_

            print("(DEBUG) All prediction probabilities:")
            for i, prob in enumerate(probabilities):
                # Convert index to gesture ID using inverse_transform
                gesture_id = int(label_encoder.inverse_transform([i])[0])
                prob_value = float(prob.item())
                all_probs[gesture_id] = prob_value
                print(f"  Gesture ID {gesture_id}: {prob_value:.4f}" + (" (PREDICTED)" if i == prediction else ""))

        return jsonify({
            "prediction": predicted_gesture_id,
            "confidence": confidence,
            "timestamp": datetime.datetime.now().isoformat(),
        })
    
    except Exception as e:
        import traceback
        traceback.print_exc()
        return jsonify({
            "error": str(e),
            "data_received": str(feature_data) if 'feature_data' in locals() else "No data"
        }), 400

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=443, debug=False)